package middleware

import (
	"bytes"
	"github.com/bighuangbee/gomod/loger"
	"github.com/gin-gonic/gin"
	//"io/ioutil"
	"net/http/httputil"
	"strings"
	"time"
)

/**
* @Author: bigHuangBee
* @Date: 2020/3/24 19:04
 */

func RequestLog() gin.HandlerFunc {

	return func(c *gin.Context) {

		startTime := time.Now()

		var requestParams interface{};
		if strings.Contains(c.Request.Header.Get("Content-Type"), "application/json"){
			body, _ := httputil.DumpRequest(c.Request, true)
			requestParams = string(body)
		} else{
			c.DefaultPostForm("test", "")
			c.Request.ParseForm()

			requestParams = c.Request.PostForm
		}

		//loger.Info(
		//	"Request " + c.Request.Method,
		//	c.ClientIP(),
		//	c.Request.RequestURI,
		//	c.Request.Header.Get("Authorization"),
		//	requestParams,
		//	c.Request.Header.Get("Content-Type"),
		//)

		writer := responeWriter{
			c.Writer,
			bytes.NewBuffer([]byte("")),
		}
		c.Writer = &writer

		c.Next()

		loger.Info(
			"\n[Request]:",
			c.ClientIP(),
			c.Request.Method,
			c.Request.RequestURI,
			time.Now().Sub(startTime),
			requestParams,
			c.Request.Header.Get("Content-Type"),
			c.Request.Header.Get("Authorization"),
			"\n[Respone]:" + writer.WriterBuff.String(),
		)
	}
}

/**
	重新实现ResponseWriter接口的Write方法
	保存请求回复的数据副本
 */
type responeWriter struct {
	gin.ResponseWriter
	WriterBuff *bytes.Buffer
}

func (r *responeWriter) Write(body []byte) (size int, err error){
	r.WriterBuff.Write(body)
	size, err = r.ResponseWriter.Write(body)	//调用ResponseWriter接口的原write
	return
}