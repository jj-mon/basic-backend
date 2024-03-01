package user

import (
	"basic-backend/internal/repository"
	"basic-backend/internal/service"
)

var _ service.UserService = (*userService)(nil)

type userService struct {
	userRepository repository.UserRepository
}

func New(userRepository repository.UserRepository) *userService {
	return &userService{
		userRepository: userRepository,
	}
}
