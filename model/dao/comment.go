package dao

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	RecordID uint
	Content  string
}

func (s *Comment) GetAllCommentByRecordID() ([]Comment, error) {
	var t []Comment
	err := DB.Model(s).Where(s).Find(&t).Error
	return t, err
}

func (s *Comment) Create() error {
	return DB.Model(s).Where(s).Create(s).Error
}
