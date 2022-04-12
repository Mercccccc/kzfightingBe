package dto

type PostComment struct {
	RecordID uint   `json:"record_id" binding:"required"`
	Content  string `json:"content" binding:"required"`
}
