package controller

import (
	"github.com/gin-gonic/gin"
	"kzfighting/model/dto"
	"kzfighting/service"
	"strconv"
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

func GetComment(c *gin.Context) {
	var req dto.GetComment
	recordID, err := strconv.Atoi(c.DefaultQuery("record_id", "0"))
	if err != nil {
		RespondError(c, service.CommitDataError)
		return
	}
	if recordID == 0 {
		RespondError(c, service.CommitDataError)
		return
	}

	req.RecordID = uint(recordID)
	code, data := service.GetComment(req)
	if code != 0 {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, data)
}
