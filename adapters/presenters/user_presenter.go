package presenters

import (
	"firestore_clean/entities"
	"firestore_clean/usecases/ports"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserPresenter struct {
	ctx echo.Context
}

func NewUserOutputPort(ctx echo.Context) ports.UserOutputPort {
	return &UserPresenter{
		ctx: ctx,
	}
}

func (u *UserPresenter) OutputUsers(users []*entities.User) error {
	return u.ctx.JSON(http.StatusOK, users)
}

func (u *UserPresenter) OutputError(err error) error {
	log.Fatal(err)
	return err
}
