package api

import "github.com/Seew0/Heal-D/domain/models"

type GetQuestionnaireResponse struct {
	ID       string `json:"id"`
	Question string `json:"question"`
	Options  []models.Option `json:"options"`
}

type GetQuestionnaireListResponse struct {
	Questionnaires []GetQuestionnaireResponse `json:"questionnaires"`
}