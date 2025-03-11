package logic

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Seew0/Heal-D/domain/api"
	"github.com/Seew0/Heal-D/domain/models"
	"github.com/Seew0/Heal-D/internal/service"
	"github.com/Seew0/Heal-D/internal/utility"
	"github.com/gin-gonic/gin"
)

type UserLogic struct {
	userService *service.UserService
}

func NewUserLogic(userService *service.UserService) *UserLogic {
	return &UserLogic{userService: userService}
}

// Create Generated User

func (l *UserLogic) CreateGenUser(c *gin.Context) {
	genedData := utility.GenerateDataEntries()
	user := &models.UserData{
		GeneratedData: genedData,
		AccountStatus: models.AccountStatusActive,
	}

	err := l.userService.CreateUser(c, *user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}

// Create User

func (l *UserLogic) CreateUser(c *gin.Context) {
	var userData api.CreateUserRequest
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var realData models.RealData
	err := json.Unmarshal([]byte(userData.RealData), &realData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid real data"})
		return
	}

	user := &models.UserData{
		RealData:      realData,
		AccountStatus: models.AccountStatusActive,
	}

	err = l.userService.CreateUser(c, *user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}

func (l *UserLogic) GetUsers(c *gin.Context) {
	users, err := l.userService.GetAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (l *UserLogic) GetUserByID(c *gin.Context) {
	id, flag := c.Params.Get("id")
	fmt.Println("ID: ", id)
	if !flag {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	user, err := l.userService.GetUserByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (l *UserLogic) GetUserScore(c *gin.Context) {
	id, flag := c.Params.Get("id")
	if !flag {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	user, err := l.userService.GetUserByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	score, err := l.userService.GetUserScore(c, user.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch score"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"score": score.Score})
}