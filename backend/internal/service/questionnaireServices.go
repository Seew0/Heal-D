package service

import (
	"context"

	"github.com/Seew0/Heal-D/domain/api"
	"github.com/Seew0/Heal-D/domain/models"
	"github.com/Seew0/Heal-D/internal/repository"
)

type QuestionnaireService struct {
	questionniarRepo *repository.QuestionnaireRepository
}

func NewQuestionnaireService(repo *repository.QuestionnaireRepository) *QuestionnaireService {
	return &QuestionnaireService{questionniarRepo: repo}
}

func (s *QuestionnaireService) GetQuestions(ctx context.Context) (api.GetQuestionnaireListResponse, error) {
	return s.questionniarRepo.GetQuestions(ctx)
}

func (s *QuestionnaireService) GetQuestionByID(ctx context.Context, id string) (api.GetQuestionnaireResponse, error) {
	return s.questionniarRepo.GetQuestionByID(ctx, id)
}

func (s *QuestionnaireService) SubmitAnswers(ctx context.Context, answers models.UserAnswer) error {
	return s.questionniarRepo.SubmitAnswers(ctx, answers)
}

func (s *QuestionnaireService) CreateScore(ctx context.Context, score models.Score) (string, error) {
	return s.questionniarRepo.CreateScore(ctx, score)
}

func (s *QuestionnaireService) GetTestStatus(ctx context.Context, id string) (bool, error) {
	return s.questionniarRepo.GetTestStatus(ctx, id)
}

func (s *QuestionnaireService) UpdateTestStatus(ctx context.Context, userID string, status string) error {
	return s.questionniarRepo.UpdateTestStatus(ctx, userID, status)
}