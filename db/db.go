package db

import (
	"GoTodoBackend/entity"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
    if err != nil {
        panic(err)
    }
	autoMigration(db)
}

func autoMigration(db *gorm.DB) {
	db.AutoMigrate(&entity.User{}, &entity.Todo{})
}

func GetDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Println("failed in getting db")
	}
    return db
}