package controller

import (
	"github.com/gin-gonic/gin"
	"kzfighting/service"
)

var HttpStatus = map[uint]int{
	service.ServerError:     500,
	service.CommitDataError: 400,
}

var Message = map[uint]string{
	service.ServerError:     "服务端错误",
	service.CommitDataError: "提交数据错误",
}

func RespondSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"data": data,
	})

	c.Abort()
	return
}

func RespondError(c *gin.Context, code uint) {
	httpStatus, ok := HttpStatus[code]
	if !ok {
		httpStatus = 403
	}
	message, ok := Message[code]
	if !ok {
		message = "未描述的错误"
	}
	c.JSON(httpStatus, map[string]interface{}{
		"code": code,
		"msg":  message,
	})
	c.Abort()
}
