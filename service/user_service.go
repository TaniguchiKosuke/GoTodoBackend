package service

import (
	"GoTodoBackend/db"
	"GoTodoBackend/entity"
	"encoding/json"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Service procides user's behavior
type Service struct{}

// User is alias of entity.User struct
type User entity.User

// GetAll is get all User
func (s Service) GetAll() ([]User, error) {
    db := db.GetDB()
    var u []User

    if err := db.Find(&u).Error; err != nil {
        return nil, err
    }

    return u, nil
}

// CreateModel is create User model
func (s *Service) RegisterUserModel(c *gin.Context) (User, error) {
    db := db.GetDB()
    var u User

    if err := c.BindJSON(&u); err != nil {
        return u, err
    }

	//Hashing the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return u, err
	}
	u.Password = string(hashedPassword)

    if err := db.Create(&u).Error; err != nil {
        return u, err
    }

    return u, nil
}

// GetByID is get a User
func (s *Service) GetUserModelByID(id string) (User, error) {
    db := db.GetDB()
    var u User

    if err := db.Where("id = ?", id).First(&u).Error; err != nil {
        return u, err
    }

    return u, nil
}

// UpdateByID is update a User
func (s *Service) UpdateUserModelByID(id string, c *gin.Context) (User, error) {
    db := db.GetDB()
    var u User

    if err := db.Where("id = ?", id).First(&u).Error; err != nil {
        return u, err
    }

    if err := c.BindJSON(&u); err != nil {
        return u, err
    }

    db.Save(&u)

    return u, nil
}

// DeleteByID is delete a User
func (s *Service) DeleteUserModelByID(id string) error {
    db := db.GetDB()
    var u User

    if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
        return err
    }

    return nil
}

type LoginInfo struct {
	Email string
	Password string
}

type SessionInfo struct {
	UserId interface{}
}

func (s *Service) LoginUserModel(c *gin.Context) error {
	db := db.GetDB()
	var loginInfo *LoginInfo
	var user User
	data, err := c.GetRawData()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &loginInfo); err != nil {
		return err
	}

	email := loginInfo.Email
	password := loginInfo.Password
	if err := db.Where("email = ?", email).Find(&user).Error; err != nil {
		return err
	}

	hashedPassword := user.Password
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return err
	}

	session := sessions.Default(c)
    session.Set("UserId", user.ID)
    session.Save()

	return nil
}

func SessionCheck() gin.HandlerFunc {
    return func(c *gin.Context) {
		var sessionInfo SessionInfo
        session := sessions.Default(c)
        sessionInfo.UserId = session.Get("UserId")

        // セッションがない場合、ログインフォームをだす
        if sessionInfo.UserId == nil {
			log.Println("Not logged in")
            c.Abort() // これがないと続けて処理されてしまう
        } else {
            c.Set("UserId", sessionInfo.UserId) // ユーザidをセット
            c.Next()
        }
    }
}

func GetRequestUserId(c *gin.Context) (User, error) {
	session := sessions.Default(c)
	requestUserId := session.Get("UserId")
	var user *User

	db := db.GetDB()
	if err := db.Find(&user, "id = ?", requestUserId).Error; err != nil {
		return *user, err
	}

	return *user, nil
}