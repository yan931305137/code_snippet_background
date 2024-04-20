package middleware

import (
	"bytes"
	"code_snippet/internal/model"
	"code_snippet/internal/types"
	"code_snippet/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// 记录请求日志
func LogRequest(c *gin.Context) {
	bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = bodyLogWriter
	//开始时间
	startTime := time.Now()
	//处理请求
	c.Next()
	// 异步记录
	go func(c *gin.Context) {
		responseBody := bodyLogWriter.body.String()
		var responseCode int
		var responseMsg string
		var responseData interface{}
		if responseBody != "" {
			response := types.JsonResult{}
			err := json.Unmarshal([]byte(responseBody), &response)
			if err == nil {
				responseCode = response.Code
				responseMsg = response.Msg
				responseData = response.Data
			}
		}
		//结束时间
		endTime := time.Now()
		// 解析令牌
		result, err := utils.ParseToken(utils.GetToken(c))
		var roleName = ""
		if result != "" {
			roleName = result
		}
		response, _ := json.Marshal(responseData)
		// 构建日志对象
		logInfo := model.Log{
			ClientIP:     c.ClientIP(),
			EdTime:       int(endTime.Sub(startTime).Milliseconds()),
			Method:       c.Request.Method,
			RequestURI:   c.Request.RequestURI,
			ResponseCode: responseCode,
			ResponseMsg:  responseMsg,
			ResponseData: string(response),
			RoleName:     roleName,
			StartTime:    startTime.Format("2006-01-02 15:04:05"),
			Level:        "info", // 可根据实际情况设置日志级别
			Msg:          "路由日志",
			Time:         endTime.Format("2006-01-02 15:04:05"),
		}
		// 将日志对象存入数据库
		_, err = utils.XormDb.Table("log").Insert(logInfo)
		if err != nil {
			fmt.Printf("日志记录失败: %v\n", err)
			return
		}
	}(c.Copy())
}
