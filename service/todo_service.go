package service

import (
	"GoTodoBackend/db"
	"GoTodoBackend/entity"
	"fmt"
)

type TodoService struct {}

type Todo entity.Todo

func (ts *TodoService) GetAllTodoModel() {
	db := db.GetDB()
	var todo []Todo
	fmt.Println(db, todo)
}