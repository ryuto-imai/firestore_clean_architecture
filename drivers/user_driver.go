package drivers

import (
	"context"
	"firestore_clean/adapters/controllers"
	"firestore_clean/adapters/gateways"
	"firestore_clean/adapters/presenters"
	"firestore_clean/drivers/database"
	"firestore_clean/usecases/interactors"
	"log"

	"github.com/labstack/echo/v4"
)

func ServeUsers(address string) {
	ctx := context.Background()
	client, err := database.GetCllient(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	user := controllers.UserController{
		OutputFactory:     presenters.NewUserOutputPort,
		InputFactory:      interactors.NewUserInputPort,
		RepositoryFactory: gateways.NewUserRepository,
		Client:            client,
	}

	e := echo.New()
	e.POST("/users", user.AddUser)
	e.GET("/users", user.GetUsers)
	e.Logger.Fatal(e.Start(address))
}
