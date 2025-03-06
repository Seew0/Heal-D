package requests

import "github.com/Seew0/Heal-D/domain/models"

type SubmitAnswersRequest struct {
	Answers []models.UserAnswer `json:"answers"`
	UserID string              `json:"user_id"`
}