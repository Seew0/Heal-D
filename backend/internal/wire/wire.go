//go:build wireinject
// +build wireinject

package wire

import (
	"os"

	"github.com/Seew0/Heal-D/internal/db"
	"github.com/Seew0/Heal-D/internal/logic"
	"github.com/Seew0/Heal-D/internal/repository"
	"github.com/Seew0/Heal-D/internal/service"
	"github.com/Seew0/Heal-D/internal/router"
	"github.com/google/wire"
)

func ProvideDB() (*db.MongoDB, error) {
	dburl := os.Getenv("DBURI")
	dbname := os.Getenv("DBNAME")
	return db.NewMongoDB(dburl, dbname)
}

// ProvideRepositories sets up all repositories.
func ProvideRepositories(db *db.MongoDB) *repository.Repository {
	return repository.NewRepository(
		repository.NewUserRepository(db),
		repository.NewQuestionnaireRepository(db),
	)
}

// ProvideServices sets up all services.
func ProvideServices(repo *repository.Repository) *service.Service {
	return service.NewService(repo)
}

// ProvideHandlers sets up all handlers.
func ProvideHandlers(svc *service.Service) *logic.Logic {
	return logic.NewLogic(svc)
}

// ProvideRouter sets up the router with handlers.
func ProvideRouters(logic *logic.Logic) *router.Router {
	return router.NewRouter(logic)
}

// InitializeApp wires everything together.
func InitializeApp() (*router.Router, error) {
	wire.Build(
		ProvideDB,
		ProvideRepositories,
		ProvideServices,
		ProvideHandlers,
		ProvideRouters,
	)
	return &router.Router{}, nil
}
