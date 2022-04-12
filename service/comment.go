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

func GetComment(req dto.GetComment) (uint, map[string]interface{}) {
	comment := &dao.Comment{
		RecordID: req.RecordID,
	}

	comments, err := comment.GetAllCommentByRecordID()
	if err != nil {
		return ServerError, nil
	}

	return SuccessCode, map[string]interface{}{
		"comments": comments,
	}
}
