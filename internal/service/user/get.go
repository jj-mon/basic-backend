package user

import (
	"basic-backend/internal/entities"
	"context"
)

func (s *userService) Get(ctx context.Context, uuid entities.UUID) (entities.User, error) {
	return s.userRepository.Get(ctx, uuid)
}
