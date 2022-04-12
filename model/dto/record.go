package dto

type PostRecord struct {
	Receiver string `json:"receiver" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

type GetRecord struct {
	Page uint
}
