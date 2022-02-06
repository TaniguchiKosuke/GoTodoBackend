package controller

import (
	"GoTodoBackend/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type TodoController struct{}

func (tc *TodoController) GetAllTodo(c *gin.Context) {
	var ts service.TodoService
	fmt.Println(ts)
}