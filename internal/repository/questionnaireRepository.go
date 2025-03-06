package repository

import (
	"context"

	"github.com/Seew0/Heal-D/domain/models"
	"github.com/Seew0/Heal-D/internal/db"
	"go.mongodb.org/mongo-driver/bson"
)

type QuestionnaireRepository struct {
	db *db.MongoDB
}

func NewQuestionnaireRepository(db *db.MongoDB) *QuestionnaireRepository {
	return &QuestionnaireRepository{db: db}
}

func (r *QuestionnaireRepository) GetQuestions(ctx context.Context) ([]models.Question, error) {
	cursor, err := r.db.QuestionsCol.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var questions []models.Question
	if err := cursor.All(ctx, &questions); err != nil {
		return nil, err
	}

	return questions, nil
}

func (r *QuestionnaireRepository) GetQuestionByID(ctx context.Context, id string) (*models.Question, error) {
	var question models.Question
	if err := r.db.QuestionsCol.FindOne(ctx, bson.M{"_id": id}).Decode(&question); err != nil {
		return nil, err
	}

	return &question, nil
}

func (r *QuestionnaireRepository) CreateScore(ctx context.Context, score models.Score) (string , error) {
	scoreID , err := r.db.ScoreUserCol.InsertOne(ctx, score)
	if err != nil {
		return "" , err
	}

	return scoreID.InsertedID.(string), nil
}

func (r *QuestionnaireRepository) SubmitAnswers(ctx context.Context, answers []models.UserAnswer) error {
	for _, answer := range answers {
		_, err := r.db.UserAnswersCol.InsertOne(ctx, answer)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *QuestionnaireRepository) UpdateTestStatus(ctx context.Context, id string, scoreID string) error {
	updatePayload := bson.M{
		"$set": bson.M{
			"testTaken": true,
			"scoreID":   scoreID,
		},
	}
	_, err := r.db.UserDataCol.UpdateByID(ctx, id, updatePayload)
	if err != nil {
		return err
	}

	return nil
}
