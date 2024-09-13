package query

import (
	"context"

	"github.com/meraf00/agora-meet/agora-be/internal/users/app/dto"
	users "github.com/meraf00/agora-meet/agora-be/internal/users/domain"
)

type UsersHandler struct {
	userRepo users.Repository
}

func NewUsersHandler(userRepo users.Repository) *UsersHandler {
	if userRepo == nil {
		panic("nil userRepo")
	}

	return &UsersHandler{
		userRepo: userRepo,
	}
}

func (h UsersHandler) Handle(ctx context.Context) ([]*dto.UserDto, error) {
	users, err := h.userRepo.FindUsers(ctx)

	if err != nil {
		return nil, err
	}

	usersResponse := dto.NewUsersDtoFromUsers(users)

	return usersResponse, nil
}
