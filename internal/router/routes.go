package router

import (
	"github.com/Seew0/Heal-D/internal/logic"
	"github.com/gin-gonic/gin"
)

type Router struct {
	userLogic *logic.UserLogic
	questionnaireLogic *logic.QuestionnaireLogic
}

func NewRouter(logic *logic.Logic) *Router {
	return &Router{
		userLogic: logic.UserLogic,
		questionnaireLogic: logic.QuestionnaireLogic,
	}
}

func (r *Router) SetupRoutes(router *gin.Engine) {
	SetupUserRoutes(router, r.userLogic)
	SetupQuestionnaireRoutes(router, r.questionnaireLogic)
}
