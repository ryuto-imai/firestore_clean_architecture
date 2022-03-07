package controllers

import (
	"context"
	"firestore_clean/entities"
	"firestore_clean/usecases/ports"

	"cloud.google.com/go/firestore"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	OutputFactory     func(ctx echo.Context) ports.UserOutputPort
	InputFactory      func(outputPort ports.UserOutputPort, repository ports.UserRepository) ports.UserInputPort
	RepositoryFactory func(client *firestore.Client) ports.UserRepository
	Client            *firestore.Client
}

func (u *UserController) AddUser(c echo.Context) error {
	user := new(entities.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	inputPort := u.newInputPort(c)

	ctx := context.Background()
	return inputPort.AddUser(ctx, user)
}

func (u *UserController) GetUsers(c echo.Context) error {
	inputPort := u.newInputPort(c)

	ctx := context.Background()
	return inputPort.GetUsers(ctx)
}

func (u *UserController) newInputPort(c echo.Context) ports.UserInputPort {
	outputPort := u.OutputFactory(c)
	repository := u.RepositoryFactory(u.Client)
	return u.InputFactory(outputPort, repository)
}
