package server

import (
	"GoTodoBackend/controller"
	"GoTodoBackend/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Init is initialize server
func Init() {
    r := router()
    r.Run()
}

func router() *gin.Engine {
    r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
    r.Use(sessions.Sessions("mysession", store))

    u := r.Group("/user")
    {
        userCtrl := controller.Controller{}
        u.GET("", userCtrl.Index)
        u.GET("/:id", userCtrl.GetUserById)
        u.POST("/signin", userCtrl.RegisterUser)
        u.PUT("/:id", userCtrl.UpdateUserById)
        u.DELETE("/:id", userCtrl.DeleteUserById)
		u.POST("/login", userCtrl.Login)
		u.Use(service.SessionCheck())
		{
			u.POST("/logout",userCtrl.Logout)
		}
    }
	
	todo := r.Group("/api")
	todo.Use(service.SessionCheck())
	{
		todoCtrl := controller.TodoController{}
		todo.GET("/todos", todoCtrl.GetAllTodo)
		todo.POST("/todo/create", todoCtrl.CreateTodo)
		todo.PUT("/todo/:id", todoCtrl.UpdateTodo)
		todo.GET("/todo/:id", todoCtrl.GetTodoById)
		todo.DELETE("/todo/:id", todoCtrl.DeleteTodoById)
	}

    return r
}