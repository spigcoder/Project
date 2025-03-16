package logService

import (
	"blog_server/core"
	"blog_server/enum"
	"blog_server/global"
	"blog_server/models"
	"blog_server/utils/jwt"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
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

// 添加content内容
func (o *OperateLog) setItem(level enum.LogLevel, label string, value any) {
	t := reflect.TypeOf(value)
	var content string
	switch t.Kind() {
	case reflect.Array, reflect.Struct, reflect.Slice:
		byteData, _ := json.Marshal(value)
		content = string(byteData)
	default:
		content = fmt.Sprintf("%v", value)
	}
	o.itmeList = append(o.itmeList, fmt.Sprintf("<div class=\"log_item %s\"><div class=\"log_item_label\">%s</div><div class=\"log_item_content\">%v</div></div>",
		level.String(),
		label, content))
}

func (o *OperateLog) SetItem(label string, value any) {
	o.setItem(enum.InfoLogLevel, label, value)
}

func (o *OperateLog) SetInfoItem(label string, value any) {
	o.setItem(enum.InfoLogLevel, label, value)
}
func (o *OperateLog) SetErrItem(label string, value any) {
	o.setItem(enum.ErrorLogLevel, label, value)
}
func (o *OperateLog) SetWarnItem(label string, value any) {
	o.setItem(enum.WarnLogLevel, label, value)
}

func (o *OperateLog) SetLink(label, herf string) {
	o.itmeList = append(o.itmeList, fmt.Sprintf("<div class=\"log_item link\"><div class=\"log_item_label\">%s</div><div class=\"log_item_content\"><a href=\"%s\" target=\"_blank\"></a></div></div>", label, herf))
}

func (o *OperateLog) SetImage(image string) {
	o.itmeList = append(o.itmeList, fmt.Sprintf("<div class=\"log_image\"><img src=\"%s\" alt=\"\"></div>", image))
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
	userID := uint(0)
	claim, err := jwt.ParseTokenByGin(o.c)
	if err == nil && claim != nil{
		userID = claim.UserID
	}

	if o.log != nil {
		// 证明更新过了
		global.DB.Model(o.log).Updates(map[string]any{
			"title": "update",
		})
		return
	}

	var newItemList []string
	if o.showRequeset {
		// TODO: 这里默认是json格式，需要根据请求头判断
		newItemList = append(newItemList, fmt.Sprintf("<div class=\"log_request\"><div class=\"log_request_head\"><span class=\"log_request_method %s\">%s</span><span class=\"log_request_path\">%s</span></div><div class=\"log_request_body\"><pre class=\"log_json_body\">%s</pre></div></div>",
			strings.ToLower(o.c.Request.Method),
			o.c.Request.Method,
			o.c.Request.URL.String(),
			string(o.request)))
	}

	//content
	newItemList = append(newItemList, o.itmeList...)

	if o.showResponse {
		newItemList = append(newItemList, fmt.Sprintf("<div class=\"log_response\"><pre class=\"log_json_body\">%s</pre></div>",
			string(o.response)))
	}

	log := models.LogModel{
		LogType: enum.OperateLogType,
		Title:   o.title,
		Content: strings.Join(newItemList, "\n"),
		Level:   o.level,
		UserID:  userID,
		IP:      ip,
		Addr:    addr,
	}

	error := global.DB.Debug().Create(&log).Error

	if error != nil {
		logrus.Warnf("保存操作日志失败: %v", error)
		return
	}
	o.log = &log
}
