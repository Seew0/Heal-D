package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client           *mongo.Client
	Database         *mongo.Database
	UserDataCol      *mongo.Collection
	ForumCol         *mongo.Collection
	MessageCol       *mongo.Collection
	ScoreMatchCol    *mongo.Collection
}

// NewMongoDB initializes a MongoDB connection.
func NewMongoDB(uri, dbName string) (*MongoDB, error) {
	if uri == "" || dbName == "" {
		return nil, fmt.Errorf("MongoDB URI or Database name is empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("MongoDB connection error: %w", err)
	}

	log.Println("🔌 Connecting to MongoDB...")
	// Ping the database to check if the connection is alive
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("MongoDB ping error: %w", err)
	}

	log.Println("✅ Connected to MongoDB!")
	db := client.Database(dbName)

	return &MongoDB{
		Client:           client,
		Database:         db,
		UserDataCol:      db.Collection("userData"),
		ForumCol:         db.Collection("forums"),
		MessageCol:       db.Collection("messages"),
		ScoreMatchCol:    db.Collection("scoreMatches"),
	}, nil
}

// Close disconnects MongoDB connection.
func (m *MongoDB) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := m.Client.Disconnect(ctx); err != nil {
		log.Printf("⚠️ Error disconnecting MongoDB: %v\n", err)
	} else {
		log.Println("✅ MongoDB connection closed.")
	}
}
