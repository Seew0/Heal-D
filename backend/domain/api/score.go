package api

import "github.com/Seew0/Heal-D/domain/models"

type SubmitAnswersRequest struct {
	Answers []models.Answer `json:"answers"`
	UserID  string    `json:"user_id"`
}
