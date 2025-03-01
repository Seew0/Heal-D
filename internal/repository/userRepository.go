package repository

import (
	"context"

	"github.com/Seew0/Heal-D/domain/models"
	"github.com/Seew0/Heal-D/internal/db"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct {
	db *db.MongoDB
}

func NewUserRepository(db *db.MongoDB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindAllUsers(ctx context.Context) ([]models.UserData, error) {
	cursor, err := r.db.UserDataCol.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []models.UserData
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}
