package repository

type Repository struct {
	UserRepository          *UserRepository
	QuestionnaireRepository *QuestionnaireRepository
}

func NewRepository(userRepo *UserRepository, questionnaireRepo *QuestionnaireRepository) *Repository {
	return &Repository{
		UserRepository: userRepo,
		QuestionnaireRepository: questionnaireRepo,
	}
}
