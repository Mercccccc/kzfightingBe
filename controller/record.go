package controller

import (
	"github.com/gin-gonic/gin"
	"kzfighting/model/dto"
	"kzfighting/service"
	"strconv"
)

func PostRecord(c *gin.Context) {
	req := &dto.PostRecord{}
	err := c.ShouldBind(req)
	if err != nil {
		RespondError(c, service.CommitDataError)
		return
	}

	code, data := service.PostRecord(req)
	if code != 0 {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, data)
	return
}

func GetRecordNumber(c *gin.Context) {
	code, data := service.GetRecordNumber()
	if code != 0 {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, data)
}

func GetRecordByPage(c *gin.Context) {
	req := dto.GetRecord{}
	number, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		RespondError(c, service.CommitDataError)
		return
	}
	if number == 0 {
		RespondError(c, service.CommitDataError)
		return
	}

	req.Page = uint(number)
	code, data := service.GetRecordByPage(req)
	if code != 0 {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, data)
}
