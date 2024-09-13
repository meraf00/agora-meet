package app

import (
	"context"

	"github.com/meraf00/agora-meet/agora-be/internal/users/app/query"
	"github.com/meraf00/agora-meet/agora-be/internal/users/infrastructure"
	"github.com/meraf00/agora-meet/agora-be/shared/config"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
}

type Queries struct {
	User  query.UserHandler
	Users query.UsersHandler
}

func NewApplication(ctx context.Context) Application {

	userRepo, err := infrastructure.NewUserMongoRepository(
		config.AppConfig.Mongodb.Uri,
		config.AppConfig.Mongodb.DbName,
		"users",
	)

	if err != nil {
		panic(err)
	}

	return Application{
		Commands: Commands{},

		Queries: Queries{
			User:  *query.NewUserHandler(userRepo),
			Users: *query.NewUsersHandler(userRepo),
		},
	}

}
