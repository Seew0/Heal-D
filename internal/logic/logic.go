package logic

import "github.com/Seew0/Heal-D/internal/service"

type Logic struct {
	UserLogic          *UserLogic
	QuestionnaireLogic *QuestionnaireLogic
}

func NewLogic(service *service.Service) *Logic {
	return &Logic{
		UserLogic: NewUserLogic(service.UserService),
		QuestionnaireLogic: NewQuestionnaireLogic(service.QuestionnaireService)}
}
