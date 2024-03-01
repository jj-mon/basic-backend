package repository

import (
	"basic-backend/internal/entities"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, uuid entities.UUID, info entities.UserInfo) error
	Get(ctx context.Context, uuid entities.UUID) (entities.User, error)
}
