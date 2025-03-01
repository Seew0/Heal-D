package logic

import (
	"net/http"

	"github.com/Seew0/Heal-D/internal/service"
	"github.com/gin-gonic/gin"
)

type UserLogic struct {
	userService *service.UserService
}

func NewUserLogic(userService *service.UserService) *UserLogic {
	return &UserLogic{userService: userService}
}

func (l *UserLogic) GetUsers(c *gin.Context) {
	users, err := l.userService.GetAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
