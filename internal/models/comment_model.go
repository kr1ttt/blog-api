package models

type Comment struct {
	ID     uint   `json:"id"`
	Text   string `json:"text" binding:"required"`
	PostID uint   `json:"post_id" binding:"required"`
	UserID uint   `json:"user_id" binding:"required"`
}
