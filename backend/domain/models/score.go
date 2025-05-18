package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Option struct {
	Text  string  `bson:"text" json:"text"`
	Score float32 `bson:"score" json:"score"`
}

type Question struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Question string             `bson:"question" validate:"required"`
	Options  []Option           `bson:"options" validate:"required,dive,required"`
}

type Answer struct {
	ID string `json:"question_id"`
	SelectedAnswer int `json:"selected_answer"`
}

type UserAnswer struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UserID     string `json:"userId"`
	Answers   []Answer           `json:"answers"`
}

type Score struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Score  float32            `bson:"score" validate:"required,min=0,max=10"`
	UserID string `bson:"userID"`
}
