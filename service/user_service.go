package service

import (
	"GoTodoBackend/db"
	"GoTodoBackend/entity"
	"encoding/json"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)

// Service procides user's behavior
type Service struct{}

// User is alias of entity.User struct
type User entity.User

// GetAll is get all User
func (s Service) GetAll() ([]User, error) {
    db := db.GetDB()
    var user []User

    if err := db.Find(&user).Error; err != nil {
        return nil, err
    }

    return user, nil
}

// CreateModel is create User model
func (s *Service) RegisterUserModel(c *gin.Context) (User, error) {
    db := db.GetDB()
    var user User

    if err := c.BindJSON(&user); err != nil {
        return user, err
    }

	//Hashing the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return user, err
	}
	user.Password = string(hashedPassword)

	u, err := uuid.NewRandom()
	if err != nil {
			return user, err
	}
	userUuid := u.String()
	user.Uuid = userUuid

    if err := db.Create(&user).Error; err != nil {
        return user, err
    }

    return user, nil
}

// GetByID is get a User
func (s *Service) GetUserModelByID(id string) (User, error) {
    db := db.GetDB()
    var user User

    if err := db.Where("uuid = ?", id).First(&user).Error; err != nil {
        return user, err
    }

    return user, nil
}

// UpdateByID is update a User
func (s *Service) UpdateUserModelByID(id string, c *gin.Context) (User, error) {
    db := db.GetDB()
    var user User

    if err := db.Where("uuid = ?", id).First(&user).Error; err != nil {
        return user, err
    }

    if err := c.BindJSON(&user); err != nil {
        return user, err
    }

    db.Save(&user)

    return user, nil
}

// DeleteByID is delete a User
func (s *Service) DeleteUserModelByID(id string) error {
    db := db.GetDB()
    var user User

    if err := db.Where("uuid = ?", id).Delete(&user).Error; err != nil {
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
    session.Set("UserId", user.Uuid)
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

func GetRequestUser(c *gin.Context) (User, error) {
	session := sessions.Default(c)
	requestUserId := session.Get("UserId")
	var user *User

	db := db.GetDB()
	if err := db.Find(&user, "uuid = ?", requestUserId).Error; err != nil {
		return *user, err
	}

	return *user, nil
}