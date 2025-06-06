package models

type Post struct {
	ID     uint   `json:"id"`
	Title  string `json:"title" binding:"required"`
	Text   string `json:"text" binding:"required"`
	UserID uint   `json:"user_id" binding:"required"`
}
