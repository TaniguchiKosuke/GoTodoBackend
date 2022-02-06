package entity

import "gorm.io/gorm"

type Todo struct {
	*gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}
