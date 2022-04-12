package controller

import (
	"github.com/gin-gonic/gin"
	"kzfighting/model/dto"
	"kzfighting/service"
)

func PostComment(c *gin.Context) {
	req := dto.PostComment{}
	err := c.ShouldBind(&req)
	if err != nil {
		RespondError(c, service.CommitDataError)
		return
	}

	code, data := service.PostComment(req)
	if code != 0 {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, data)
}
