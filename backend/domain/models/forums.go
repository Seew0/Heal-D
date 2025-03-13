package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Forum struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	Name        string               `bson:"name" validate:"required"`
	Description string               `bson:"description,omitempty"`
	MinScore    int                  `bson:"minScore"`
	MaxScore    int                  `bson:"maxScore"`
	CreatedAt   time.Time            `bson:"createdAt"`
	Users       []primitive.ObjectID `bson:"users,omitempty"` // References to UserData
}

type Message struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ForumID   primitive.ObjectID `bson:"forumID"` // Foreign key to Forum
	UserID    primitive.ObjectID `bson:"userID"`  // Foreign key to UserData
	Content   string             `bson:"content" validate:"required"`
	CreatedAt time.Time          `bson:"createdAt"`
}

type ScoreMatch struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	UserID  primitive.ObjectID `bson:"userID"` // Foreign key to UserData
	Score   int                `bson:"score"`
	ForumID primitive.ObjectID `bson:"forumID"` // Foreign key to Forum
}
