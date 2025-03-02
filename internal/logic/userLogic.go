package logic

import (
	"net/http"

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

func (l *UserLogic) CreateUser(c *gin.Context) {
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
