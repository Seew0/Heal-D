package repository

type Repository struct {
	UserRepository *UserRepository
	// add more
}

func NewRepository(userRepo *UserRepository) *Repository {
	return &Repository{UserRepository: userRepo}
}
