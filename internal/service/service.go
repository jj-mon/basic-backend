package service

import (
	"basic-backend/internal/entities"
	"context"
)

type UserService interface {
	Create(ctx context.Context, info entities.UserInfo) (entities.UUID, error)
	Get(ctx context.Context, uuid entities.UUID) (entities.User, error)
}
