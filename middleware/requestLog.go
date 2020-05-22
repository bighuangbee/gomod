package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/bighuangbee/gomod/loger"
	"github.com/gin-gonic/gin"
	"io/ioutil"
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

		var requestParams string;
		if c.Request.Method == "GET"{
			if len(c.Request.URL.Query()) > 0{
				params, _ := json.Marshal(c.Request.URL.Query())
				requestParams = string(params)
			}
		}else{
			if strings.Contains(c.Request.Header.Get("Content-Type"), "application/json"){
				body,_ := ioutil.ReadAll(c.Request.Body)
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
				requestParams = string(body)
			} else{
				c.DefaultPostForm("default", "")
				c.Request.ParseForm()

				result,_ := json.Marshal(c.Request.PostForm)
				requestParams = string(result)
			}
		}

		//writer := responeWriter{
		//	c.Writer,
		//	bytes.NewBuffer([]byte("")),
		//}
		//c.Writer = &writer

		c.Next()

		loger.Info(
			"[Request]:",
			c.ClientIP(),
			c.Request.Method,
			c.Request.RequestURI,
			time.Now().Sub(startTime),
			requestParams,
			c.Request.Header.Get("Content-Type"),
			c.Request.Header.Get("Authorization"),
			//"\n[Respone]:" + writer.WriterBuff.String(),
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