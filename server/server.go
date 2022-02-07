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

    u := r.Group("/user")
    {
        ctrl := controller.Controller{}
        u.GET("", ctrl.Index)
        u.GET("/:id", ctrl.GetUserById)
        u.POST("", ctrl.RegisterUser)
        u.PUT("/:id", ctrl.UpdateUserById)
        u.DELETE("/:id", ctrl.DeleteUserById)
    }

	todo := r.Group("/api")
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