package server

import (
	"GoTodoBackend/controller"

	"github.com/gin-gonic/gin"
)

// Init is initialize server
func Init() {
    r := router()
    r.Run()
}

func router() *gin.Engine {
    r := gin.Default()

    u := r.Group("/users")
    {
        ctrl := controller.Controller{}
        u.GET("", ctrl.Index)
        u.GET("/:id", ctrl.Show)
        u.POST("", ctrl.Create)
        u.PUT("/:id", ctrl.Update)
        u.DELETE("/:id", ctrl.Delete)
    }

	todo := r.Group("/todo")
	{
		ctrl := controller.TodoController{}
		todo.GET("", ctrl.GetAllTodo)
		todo.POST("", ctrl.CreateTodo)
		todo.PUT("/:id", ctrl.UpdateTodo)
		todo.GET("/:id", ctrl.GetTodoById)
		todo.DELETE("/:id", ctrl.DeleteTodoById)
	}

    return r
}