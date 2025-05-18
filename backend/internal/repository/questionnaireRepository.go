package repository

import (
	"context"

	"github.com/Seew0/Heal-D/domain/api"
	"github.com/Seew0/Heal-D/domain/models"
	"github.com/Seew0/Heal-D/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuestionnaireRepository struct {
	db *db.MongoDB
}

func NewQuestionnaireRepository(db *db.MongoDB) *QuestionnaireRepository {
	return &QuestionnaireRepository{db: db}
}

func (r *QuestionnaireRepository) GetQuestions(ctx context.Context) (api.GetQuestionnaireListResponse, error) {
	cursor, err := r.db.QuestionsCol.Find(ctx, bson.M{})
	if err != nil {
		return api.GetQuestionnaireListResponse{}, err
	}
	defer cursor.Close(ctx)

	var questionDBResponse []models.Question
	if err := cursor.All(ctx, &questionDBResponse); err != nil {
		return api.GetQuestionnaireListResponse{}, err
	}
	var questions api.GetQuestionnaireListResponse
	for _, q := range questionDBResponse {
		questions.Questionnaires = append(questions.Questionnaires, api.GetQuestionnaireResponse{
			Question: q.Question,
			Options:  q.Options,
			ID:       q.ID.Hex(),
		})
	}

	return questions, nil
}

func (r *QuestionnaireRepository) GetQuestionByID(ctx context.Context, id string) (api.GetQuestionnaireResponse, error) {
	var questionDBResponse models.Question
	questionID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return api.GetQuestionnaireResponse{}, err
	}

	if err := r.db.QuestionsCol.FindOne(ctx, bson.M{"_id": questionID}).Decode(&questionDBResponse); err != nil {
		return api.GetQuestionnaireResponse{}, err
	}

	question := &api.GetQuestionnaireResponse{
		Question: questionDBResponse.Question,
		Options:  questionDBResponse.Options,
		ID:       questionDBResponse.ID.Hex(),
	}

	return *question, nil
}

func (r *QuestionnaireRepository) CreateScore(ctx context.Context, score models.Score) (string, error) {
	scoreID, err := r.db.ScoreUserCol.InsertOne(ctx, score)
	if err != nil {
		return "", err
	}

	scoreIDString := scoreID.InsertedID.(primitive.ObjectID).Hex()

	return scoreIDString, nil
}

func (r *QuestionnaireRepository) SubmitAnswers(ctx context.Context, answers models.UserAnswer) error {
	_, err := r.db.UserAnswersCol.InsertOne(ctx, answers)
	if err != nil {
		return err
	}
	return nil
}

func (r *QuestionnaireRepository) GetTestStatus(ctx context.Context, id string) (bool, error) {
	var user models.UserData
	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	if err := r.db.UserDataCol.FindOne(ctx, bson.M{"_id": mongoID}).Decode(&user); err != nil {
		return false, err
	}

	return user.TestTaken, nil
}

func (r *QuestionnaireRepository) UpdateTestStatus(ctx context.Context, id string, scoreID string) error {
	updatePayload := bson.M{
		"$set": bson.M{
			"testTaken": true,
			"scoreID":   scoreID,
		},
	}

	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.db.UserDataCol.UpdateByID(ctx, userID, updatePayload)
	if err != nil {
		return err
	}

	return nil
}
