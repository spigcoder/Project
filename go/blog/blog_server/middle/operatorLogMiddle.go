package middle

import (
	logService "blog_server/service/log_service"

	"github.com/gin-gonic/gin"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	Body []byte
}

func (w *responseBodyWriter) Write(b []byte) (int, error) {
	w.Body = b
	return w.ResponseWriter.Write(b)
}

func OperatorLogMiddle(c *gin.Context) {
	//请求中间件
	log := logService.NewOperateLog(c)
	log.SetRequest(c)
	//目的是后面的视图与当前使用同一个log
	c.Set("log", log)

	res := &responseBodyWriter{ResponseWriter: c.Writer}
	c.Writer = res
	c.Next()
	//相应中间件
	log.SetResponse(res.Body)
	log.Save()
}
