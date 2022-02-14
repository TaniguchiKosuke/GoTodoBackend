package service

import (
	"GoTodoBackend/db"
	"GoTodoBackend/entity"

	"github.com/gin-gonic/gin"
)

type TodoService struct {}

type Todo entity.Todo

func (ts *TodoService) GetAllTodoModel(c *gin.Context) ([]Todo, error) {
	db := db.GetDB()
	var todo []Todo
	requestUser, err := GetRequestUser(c)
	if err != nil {
		return todo, err
	}
	
	if err := db.Where("user_id = ?", requestUser.Uuid).Find(&todo).Error; err != nil {
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

	requestUser, err := GetRequestUser(c)
	if err != nil {
		return todo, err
	}

	todo.UserID = requestUser.Uuid
	db.Save(&todo)

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