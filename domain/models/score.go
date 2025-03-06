package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Option struct {
	Text  string  `bson:"text" json:"text"`
	Value float32 `bson:"value" json:"value"`
}

type Question struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Question string             `bson:"question" validate:"required"`
	Options  []Option           `bson:"options" validate:"required,dive,required"`
}

type UserAnswer struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UserID     primitive.ObjectID `json:"userId"`
	QuestionID primitive.ObjectID `json:"questionId"`
	Selected   int                `json:"selected"`
}
type Score struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Score  float32            `bson:"score" validate:"required,min=0,max=10"`
	UserID primitive.ObjectID `bson:"userID"`
}
