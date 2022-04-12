package dao

import "gorm.io/gorm"

type Record struct {
	gorm.Model
	Content  string
	Receiver string
}

func (s *Record) Retrieve() error {
	return DB.Model(s).Where(s).Last(s).Error
}

func (s *Record) Create() error {
	return DB.Model(s).Where(s).Create(s).Error
}

func GetRecordNumber() (int64, error) {
	var number int64
	err := DB.Model(&Record{}).Table("records").Count(&number).Error
	return number, err
}
