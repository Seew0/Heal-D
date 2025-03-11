package logic

import (
	"fmt"
	"net/http"

	requests "github.com/Seew0/Heal-D/domain/api"
	"github.com/Seew0/Heal-D/domain/models"
	"github.com/Seew0/Heal-D/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuestionnaireLogic struct {
	questionnaireService *service.QuestionnaireService
}

func NewQuestionnaireLogic(questionnaireService *service.QuestionnaireService) *QuestionnaireLogic {
	return &QuestionnaireLogic{questionnaireService: questionnaireService}
}

func (l *QuestionnaireLogic) GetQuestions(c *gin.Context) {
	questions, err := l.questionnaireService.GetQuestions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch questions"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"questions": questions})
}

func (l *QuestionnaireLogic) SubmitAnswers(c *gin.Context) {
	var answersReq requests.SubmitAnswersRequest
	if err := c.BindJSON(&answersReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Getting answers
	answers := answersReq.Answers

	// Calculate total score
	var totalScore float32
	for _, answer := range answers {
		fmt.Println(answer.QuestionID.Hex())
		question, err := l.questionnaireService.GetQuestionByID(c, answer.QuestionID.Hex())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch question"})
			return
		}
		answer.UserID, err = primitive.ObjectIDFromHex(answersReq.UserID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}
		totalScore += question.Options[answer.Selected-1].Score
	}

	totalScore = totalScore/10
	fmt.Println("Total Score: ", totalScore)
	// Updating answers in DB
	var userAnswersData []models.UserAnswer
	for _, answer := range answers {
		answerUserID, err := primitive.ObjectIDFromHex(answersReq.UserID)

		userAnswersData = append(userAnswersData, models.UserAnswer{
			UserID:     answerUserID,
			QuestionID: answer.QuestionID,
			Selected:   answer.Selected,
		})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}
	}

	err := l.questionnaireService.SubmitAnswers(c, userAnswersData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit answers"})
		return
	}

	// UPDATE SCORE

	var score models.Score
	score.Score = totalScore
	userID, err := primitive.ObjectIDFromHex(answersReq.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	score.UserID = userID

	var scoreID string
	scoreID, err = l.questionnaireService.CreateScore(c, score)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create score"})
		return
	}

	// Update USER

	err = l.questionnaireService.UpdateTestStatus(c, userID.Hex(), scoreID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Answers submitted"})
}
