package models

import (
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Emails string

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func (e Emails) CheckValid() bool {
	email := string(e)
	return emailRegex.MatchString(email)
}

type RealData struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name" validate:"required,min=3,max=50"`
	Age      int                `bson:"age" validate:"required,min=1,max=100"`
	Location string             `bson:"location" validate:"required,min=3,max=50"`
	Email    Emails             `bson:"email" validate:"required,email"`
}

type GeneratedData struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	GenedName     string             `bson:"genedName" validate:"required,min=3,max=50"`
	Age           int                `bson:"age" validate:"required,min=1,max=100"`
	GenedPassword string             `bson:"genedPassword" validate:"required,min=3,max=50"`
	GenedLocation string             `bson:"genedLocation" validate:"required,min=3,max=50"`
}

type UserData struct {
	ID            primitive.ObjectID     `bson:"_id,omitempty"`
	GeneratedData GeneratedData          `bson:"generatedData" validate:"required"`
	CreatedAt     time.Time              `bson:"createdAt"`
	LastActiveAt  time.Time              `bson:"lastActiveAt"`
	AccountStatus string                 `bson:"accountStatus" validate:"oneof=active banned"`
	Score         float32                `bson:"score" validate:"required,min=0,max=10"`
	Preferences   map[string]interface{} `bson:"preferences"`
}
