package controller

import (
	"GoTodoBackend/service"
	"log"

	"github.com/gin-gonic/gin"
)

type TodoController struct{}

func (tc *TodoController) GetAllTodo(c *gin.Context) {
	var ts service.TodoService
	todo, err := ts.GetAllTodoModel()
	
	if err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
		return
	}

	c.JSON(200, todo)
}

func (tc *TodoController) CreateTodo(c *gin.Context) {
	var ts service.TodoService
	todo, err := ts.CreateTodoModel(c)
	if err != nil {
		c.AbortWithStatus(400)
		log.Println(err)
		return
	}

	c.JSON(201, todo)
}

func (tc *TodoController) UpdateTodo(c *gin.Context) {
	var ts service.TodoService
	id := c.Param("id")
	todo, err := ts.GetTodoModelById(c, id)

	if err != nil {
		c.AbortWithStatus(400)
		log.Println(err)
		return
	}
	c.JSON(200, todo)
}