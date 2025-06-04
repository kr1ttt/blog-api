package models

type User struct {
	User_id       int    `json:"id" binding:"required"`
	Login         string `json:"login" binding:"required"`
	User_password string `json:"password" binding:"required"`
}
