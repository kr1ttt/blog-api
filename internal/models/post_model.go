package models

type Post struct {
	ID     uint   `json:"id" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Text   string `json:"text" binding:"required"`
	UserID uint   `json:"userID" binding:"required"`
}
