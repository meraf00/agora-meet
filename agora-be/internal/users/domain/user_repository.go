package users

import (
	"context"

	"github.com/meraf00/agora-meet/agora-be/internal/users/domain/entities"
)

type Repository interface {
	FindUser(ctx context.Context, id string) (*entities.User, error)
	SaveUser(ctx context.Context, user entities.User) error
	FindUsers(ctx context.Context) ([]*entities.User, error)
}
