package service

import (
	"context"

	"github.com/Seew0/Heal-D/domain/models"
	"github.com/Seew0/Heal-D/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user models.UserData) error {
	return s.userRepo.CreateUser(ctx, user)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]models.UserData, error) {
	return s.userRepo.FindAllUsers(ctx)
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*models.UserData, error) {
	return s.userRepo.FindUserByID(ctx, id)
}

func (s *UserService) GetUserScore(ctx context.Context, id string) (*models.Score, error) {
	return s.userRepo.GetUserScore(ctx, id)
}