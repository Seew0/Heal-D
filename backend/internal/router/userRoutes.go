package router

import (
	"github.com/Seew0/Heal-D/internal/logic"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, userLogic *logic.UserLogic) {
	userGroup := r.Group("/users")
	{
		userGroup.GET("/createGenUser", userLogic.CreateGenUser)
		userGroup.POST("/createUser", userLogic.CreateUser)
		userGroup.GET("/getUserById/:id",userLogic.GetUserByID)
		userGroup.GET("/getAllUsers", userLogic.GetUsers)
		userGroup.GET("/getUserScore/:id", userLogic.GetUserScore)
	}
}
