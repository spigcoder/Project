package logService

import (
	"blog_server/core"
	"blog_server/enum"
	"blog_server/global"
	"blog_server/models"
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type OperateLog struct {
	c            *gin.Context
	title        string
	level        enum.LogLevel
	request      []byte
	response     []byte
	log          *models.LogModel
	showRequeset bool
	showResponse bool
	itmeList     []string
}

func (o *OperateLog) ShowRequest() {
	o.showRequeset = true
}

func (o *OperateLog) ShowResponse() {
	o.showResponse = true
}

func (o *OperateLog) SetTitle(title string) {
	o.title = title
}

func (o *OperateLog) SetLevel(level enum.LogLevel) {
	o.level = level
}

func (o *OperateLog) SetRequest(c *gin.Context) {
	byteData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Error(err)
	}
	c.Request.Body = io.NopCloser(bytes.NewReader(byteData))
	o.request = byteData
}

func (o *OperateLog) SetResponse(response []byte) {
	o.response = response
}

func GetLogByGin(c *gin.Context) *OperateLog {
	_log, ok := c.Get("log")
	if !ok {
		fmt.Println("log not found")
		return NewOperateLog(c)
	}
	log, ok := _log.(*OperateLog)
	if !ok {
		fmt.Println("log type error")
		return NewOperateLog(c)
	}
	return log
}

func NewOperateLog(c *gin.Context) *OperateLog {
	return &OperateLog{
		c: c,
	}
}

func (o *OperateLog) Save() {
	ip := o.c.ClientIP()
	addr := core.GetIpAddr(ip)
	userId := uint(1)

	if o.log != nil {
		// 证明更新过了
		global.DB.Model(o.log).Updates(map[string]any{
			"title": "update",
		})
		return
	}

	if o.showRequeset {
		// TODO: 这里默认是json格式，需要根据请求头判断
		o.itmeList = append(o.itmeList, fmt.Sprintf("<div class=\"log_request\"><div class=\"log_request_head\"><span class=\"log_request_method delete\">%s</span><span class=\"log_request_path\">%s</span></div><div class=\"log_request_body\"><pre class=\"log_json_body\">%s</pre></div></div>",
			o.c.Request.Method,
			o.c.Request.URL.String(),
			string(o.request)))
	}

	if o.showResponse {
		o.itmeList = append(o.itmeList, fmt.Sprintf("<div class=\"log_response\"><pre class=\"log_json_body\">%s</pre></div>",
			string(o.response)))
	}

	log := models.LogModel{
		LogType: enum.OperateLogType,
		Title:   o.title,
		Content: strings.Join(o.itmeList, "\n"),
		Level:   o.level,
		UserID:  userId,
		IP:      ip,
		Addr:    addr,
	}

	error := global.DB.Create(&log).Error

	if error != nil {
		logrus.Warnf("保存操作日志失败: %v", error)
		return
	}
	o.log = &log
}
