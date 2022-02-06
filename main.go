package main

import (
	"GoTodoBackend/db"
	"GoTodoBackend/server"

	"github.com/gin-gonic/gin"
)



func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.String(200, "Hello, World")
    })
    db.Init()

    server.Init()
}