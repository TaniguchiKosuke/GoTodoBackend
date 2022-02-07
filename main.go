package main

import (
	"GoTodoBackend/db"
	"GoTodoBackend/server"
)



func main() {
    db.Init()
    server.Init()
}