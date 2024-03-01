package user

import (
	"basic-backend/internal/entities"
	"context"
	"github.com/google/uuid"
)

func (s *userService) Create(ctx context.Context, info entities.UserInfo) (entities.UUID, error) {
	var id entities.UUID

	uuid, err := uuid.NewUUID()
	if err != nil {
		return id, err
	}

	id = entities.UUID(uuid.String())

	return id, nil
}
