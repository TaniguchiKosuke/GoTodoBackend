package service

import (
	"GoTodoBackend/db"
	"GoTodoBackend/entity"

	"github.com/gin-gonic/gin"
)

type TodoService struct {}

type Todo entity.Todo

func (ts *TodoService) GetAllTodoModel() ([]Todo, error) {
	db := db.GetDB()
	var todo []Todo
	
	if err := db.Find(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}

func (ts *TodoService) CreateTodoModel(c *gin.Context) (Todo, error) {
	db := db.GetDB()
	var todo Todo

	if err := c.BindJSON(&todo); err != nil {
		return todo, err
	}

	if err := db.Create(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}

func (ts *TodoService) UpdateTodoModelById(c *gin.Context, id string) (Todo, error) {
	db := db.GetDB()
	var todo Todo

	if err := db.Where("id = ?", id).Find(&todo).Error; err != nil {
		return todo, err
	}

	if err := c.BindJSON(&todo); err != nil {
		return todo, err
	}

	db.Save(&todo)
	return todo, nil
}

func (ts *TodoService) GetTodoModelById(c *gin.Context, id string) (Todo, error) {
	db := db.GetDB()
	var todo Todo

	if err := db.Where("id = ?", id).Find(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}

func (ts *TodoService) DeleteTodoModelById(c *gin.Context, id string) error {
	db := db.GetDB()
	var todo Todo

	if err := db.Where("id = ?", id).Delete(&todo).Error; err != nil {
		return err
	}

	return nil
}