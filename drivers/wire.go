//go:build wireinject
// +build wireinject

package drivers

import (
	"context"
	"firestore_clean/adapters/controllers"
	"firestore_clean/adapters/gateways"
	"firestore_clean/adapters/presenters"
	"firestore_clean/drivers/database"
	"firestore_clean/usecases/interactors"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func InitializeUserDriver(ctx context.Context) (User, error) {
	wire.Build(database.NewCllient, echo.New, NewOutputFactory, NewInputFactory, NewRepositoryFactory, controllers.NewUserController, NewUserDriver)
	return &UserDriver{}, nil
}

func NewOutputFactory() controllers.OutputFactory {
	return presenters.NewUserOutputPort
}

func NewInputFactory() controllers.InputFactory {
	return interactors.NewUserInputPort
}

func NewRepositoryFactory() controllers.RepositoryFactory {
	return gateways.NewUserRepository
}
