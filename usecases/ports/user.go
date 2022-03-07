package ports

import (
	"context"
	"firestore_clean/entities"
)

type UserInputPort interface {
	AddUser(ctx context.Context, user *entities.User) error
	GetUsers(ctx context.Context) error
}

type UserOutputPort interface {
	OutputUsers([]*entities.User) error
	OutputError(error) error
}

type UserRepository interface {
	AddUser(ctx context.Context, user *entities.User) ([]*entities.User, error)
	GetUsers(ctx context.Context) ([]*entities.User, error)
}
