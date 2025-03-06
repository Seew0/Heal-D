package router

import (
	"github.com/Seew0/Heal-D/internal/logic"
	"github.com/gin-gonic/gin"
)

func SetupQuestionnaireRoutes(r *gin.Engine, questionnaireLogic *logic.QuestionnaireLogic) {
	questionnaireRoutes := r.Group("/questionnaire")
	{
		questionnaireRoutes.GET("/getQuestions", questionnaireLogic.GetQuestions)
		questionnaireRoutes.POST("/submitAnswers", questionnaireLogic.SubmitAnswers)
	}
}
