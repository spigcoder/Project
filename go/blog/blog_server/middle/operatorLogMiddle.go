package middle

import (
	"bytes"
	"fmt"
	"io"

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
	byteData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("byteData:", string(byteData))
	c.Request.Body = io.NopCloser(bytes.NewReader(byteData))
	rBW := &responseBodyWriter{ResponseWriter: c.Writer}
	c.Writer = rBW
	c.Next()
	//相应中间件
	fmt.Println(string(rBW.Body))
}
