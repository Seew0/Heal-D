package utility

import (
	"math/rand"
	"time"

	"github.com/Seew0/Heal-D/domain/models"
	"github.com/brianvoe/gofakeit/v6"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateDataEntries() models.GeneratedData {
	var data models.GeneratedData
	rand.Seed(time.Now().UnixNano())

	data = models.GeneratedData{
		ID:            primitive.NewObjectID(),
		GenedName:     gofakeit.Name(),
		Age:           rand.Intn(100) + 1,
		GenedPassword: gofakeit.Password(true, true, true, true, false, 10),
	}

	return data
}
