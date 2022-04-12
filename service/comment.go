package service

import (
	"kzfighting/model/dao"
	"kzfighting/model/dto"
)

func PostComment(req dto.PostComment) (uint, map[string]interface{}) {
	comment := &dao.Comment{
		Content:  req.Content,
		RecordID: req.RecordID,
	}

	err := comment.Create()
	if err != nil {
		return ServerError, nil
	}

	return SuccessCode, nil
}
