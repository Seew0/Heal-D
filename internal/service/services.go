package service

import "github.com/Seew0/Heal-D/internal/repository"

type Service struct {
	UserService *UserService
	// add more
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserService: NewUserService(repo.UserRepository),
	}
}
