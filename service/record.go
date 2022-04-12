package service

import (
	"kzfighting/model/dao"
	"kzfighting/model/dto"
)

func PostRecord(req *dto.PostRecord) (uint, map[string]interface{}) {
	record := dao.Record{
		Receiver: req.Receiver,
		Content:  req.Content,
	}

	err := record.Create()
	if err != nil {
		return ServerError, nil
	}

	return SuccessCode, nil
}

func GetRecordNumber() (uint, map[string]interface{}) {
	number, err := dao.GetRecordNumber()
	if err != nil {
		return ServerError, nil
	}

	return SuccessCode, map[string]interface{}{
		"number": number,
	}
}
