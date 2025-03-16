package logApi

import (
	logQuery "blog_server/common/log_query"
	"blog_server/common/res"
	"blog_server/enum"
	"blog_server/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

type LogApi struct {
}

type LogListRequest struct {
	logQuery.PageInfo
	Level       enum.LogLevel `form:"level"`  // 日志级别 1 2 3
	UserID      uint          `form:"userID"` // 用户id
	IP          string        `form:"ip"`
	LoginStatus bool          `form:"loginStatus"` // 是否登录
	UserName    string        `form:"userName"`    // 用户名
	LogType     enum.LogType  `form:"logType"`     // 日志类型 1 2 3
}

type ListResponse struct {
	models.LogModel
	UserName   string `json:"userName"`
	UserAvatar string `json:"userAvatar"`
}

func (l LogApi) LogListView(c *gin.Context) {
	var cr LogListRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		fmt.Println(err)
		res.FailWithError(err, c)
		return
	}
	model := models.LogModel{
		LogType:     cr.LogType,
		Level:       cr.Level,
		LoginStatus: cr.LoginStatus,
		IP:          cr.IP,
		UserName:    cr.UserName,
		UserID:      cr.UserID,
	}
	list, count, err := logQuery.ListQuery[models.LogModel](model, logQuery.Options{
		PageInfo:     cr.PageInfo,
		Likes:        []string{"title"},
		Preloads:     []string{"UserModel"},
		DefaultOrder: "created_at desc",
		Debug:        true,
	})
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	_list := make([]ListResponse, 0)
	for _, model := range list {
		_list = append(_list, ListResponse{
			LogModel:   model,
			UserName:   model.UserModel.Username,
			UserAvatar: model.UserModel.Avatar,
		})
	}
	res.OkWithList(_list, count, c)
}
