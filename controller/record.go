package controller

import (
	"github.com/gin-gonic/gin"
	"kzfighting/model/dto"
	"kzfighting/service"
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
