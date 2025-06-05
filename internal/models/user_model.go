package models

type User struct {
	UserID       uint   `json:"id" binding:"required"`
	Login        string `json:"login" binding:"required"`
	UserPassword string `json:"password" binding:"required"`
}
