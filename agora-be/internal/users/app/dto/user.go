package dto

import "github.com/meraf00/agora-meet/agora-be/internal/users/domain/entities"

type UserDto struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserDtoFromUser(user entities.User) *UserDto {
	return &UserDto{
		Name:  user.Name,
		Email: user.Email,
		Id:    user.Id,
	}
}

func NewUsersDtoFromUsers(users []*entities.User) []*UserDto {
	var result []*UserDto

	for _, user := range users {
		result = append(result, &UserDto{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return result
}
