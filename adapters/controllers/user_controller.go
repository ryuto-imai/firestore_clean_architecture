package controllers

import (
	"context"
	"firestore_clean/entities"
	"firestore_clean/usecases/ports"

	"cloud.google.com/go/firestore"
	"github.com/labstack/echo/v4"
)

type User interface {
	AddUser(ctx context.Context) func(c echo.Context) error
	GetUsers(ctx context.Context) func(c echo.Context) error
}

type OutputFactory func(echo.Context) ports.UserOutputPort
type InputFactory func(ports.UserOutputPort, ports.UserRepository) ports.UserInputPort
type RepositoryFactory func(*firestore.Client) ports.UserRepository

type UserController struct {
	outputFactory     OutputFactory
	inputFactory      InputFactory
	repositoryFactory RepositoryFactory
	client            *firestore.Client
}

func NewUserController(outputFactory OutputFactory, inputFactory InputFactory, repositoryFactory RepositoryFactory, client *firestore.Client) User {
	return &UserController{
		outputFactory:     outputFactory,
		inputFactory:      inputFactory,
		repositoryFactory: repositoryFactory,
		client:            client,
	}
}

func (u *UserController) AddUser(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		user := new(entities.User)
		if err := c.Bind(user); err != nil {
			return err
		}

		return u.newInputPort(c).AddUser(ctx, user)
	}
}

func (u *UserController) GetUsers(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		return u.newInputPort(c).GetUsers(ctx)
	}
}

func (u *UserController) newInputPort(c echo.Context) ports.UserInputPort {
	outputPort := u.outputFactory(c)
	repository := u.repositoryFactory(u.client)
	return u.inputFactory(outputPort, repository)
}
