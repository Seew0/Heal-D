// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/Seew0/Heal-D/internal/db"
	"github.com/Seew0/Heal-D/internal/logic"
	"github.com/Seew0/Heal-D/internal/repository"
	"github.com/Seew0/Heal-D/internal/router"
	"github.com/Seew0/Heal-D/internal/service"
	"os"
)

// Injectors from wire.go:

// InitializeApp wires everything together.
func InitializeApp() (*router.Router, error) {
	mongoDB, err := ProvideDB()
	if err != nil {
		return nil, err
	}
	repository := ProvideRepositories(mongoDB)
	service := ProvideServices(repository)
	logic := ProvideHandlers(service)
	routerRouter := ProvideRouters(logic)
	return routerRouter, nil
}

// wire.go:

func ProvideDB() (*db.MongoDB, error) {
	dburl := os.Getenv("DBURI")
	dbname := os.Getenv("DBNAME")
	return db.NewMongoDB(dburl, dbname)
}

// ProvideRepositories sets up all repositories.
func ProvideRepositories(db2 *db.MongoDB) *repository.Repository {
	return repository.NewRepository(repository.NewUserRepository(db2))
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
func ProvideRouters(logic2 *logic.Logic) *router.Router {
	return router.NewRouter(logic2)
}
