package controller

import (
	"GoTodoBackend/service"
	"fmt"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Controller is user controlller
type Controller struct{}

// Index action: GET /users
func (pc *Controller) Index(c *gin.Context) {
    var s service.Service
    p, err := s.GetAll()

    if err != nil {
        c.AbortWithStatus(404)
        return
    }
	c.JSON(200, p)
}

// Create action: POST /users
func (pc *Controller) RegisterUser(c *gin.Context) {
    var s service.Service
    p, err := s.RegisterUserModel(c)

    if err != nil {
        c.AbortWithStatus(400)
        fmt.Println(err)
		return
    }
	c.JSON(201, p)
}

// Show action: GET /users/:id
func (pc *Controller) GetUserById(c *gin.Context) {
    id := c.Params.ByName("id")
    var s service.Service
    p, err := s.GetUserModelByID(id)

    if err != nil {
        c.AbortWithStatus(404)
        fmt.Println(err)
		return
    }
	c.JSON(200, p)
}

// Update action: PUT /users/:id
func (pc *Controller) UpdateUserById(c *gin.Context) {
    id := c.Params.ByName("id")
    var s service.Service
    p, err := s.UpdateUserModelByID(id, c)

    if err != nil {
        c.AbortWithStatus(400)
        fmt.Println(err)
    }
    c.JSON(200, p)
}

// Delete action: DELETE /users/:id
func (pc *Controller) DeleteUserById(c *gin.Context) {
    id := c.Params.ByName("id")
    var s service.Service

    if err := s.DeleteUserModelByID(id); err != nil {
        c.AbortWithStatus(403)
        fmt.Println(err)
		return
    }
    c.JSON(204, gin.H{"id #" + id: "deleted"})
}

func (pc *Controller) Login(c *gin.Context) {
	var s service.Service
	
	if err := s.LoginUserModel(c); err != nil {
		c.AbortWithStatus(401)
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{"status": "Login successed"})
}

func (pc *Controller) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(200, gin.H{"status": "successed"})
}