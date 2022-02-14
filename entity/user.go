package entity

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Uuid     string `json:"uuid" gorm:"primaryKey"`
    Username string `json:"username"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}