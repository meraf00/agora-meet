package query

import (
	"context"

	"github.com/meraf00/agora-meet/agora-be/internal/users/app/dto"
	users "github.com/meraf00/agora-meet/agora-be/internal/users/domain"
)

type User struct {
	Id string
}

type UserHandler struct {
	userRepo users.Repository
}

func NewUserHandler(userRepo users.Repository) *UserHandler {
	if userRepo == nil {
		panic("nil userRepo")
	}

	return &UserHandler{
		userRepo: userRepo,
	}
}

func (h UserHandler) Handle(ctx context.Context, query User) (*dto.UserDto, error) {
	user, err := h.userRepo.FindUser(ctx, query.Id)

	if err != nil {
		return nil, err
	}

	userResponse := dto.NewUserDtoFromUser(*user)

	return userResponse, nil
}
