package entity

import "gorm.io/gorm"

type Todo struct {
	*gorm.Model
	UserID  string `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
