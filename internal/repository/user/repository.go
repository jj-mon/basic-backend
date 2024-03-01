package user

import (
	"basic-backend/internal/entities"
	"basic-backend/internal/repository"
	"context"
	"errors"
	"fmt"
	"time"
)

var _ repository.UserRepository = (*userRepository)(nil)

type userRepository struct {
	userTable map[entities.UUID]entities.User
}

func New() *userRepository {
	r := &userRepository{userTable: make(map[entities.UUID]entities.User)}
	r.userTable["123"] = entities.User{
		UUID: "123",
		Info: entities.UserInfo{
			Login:    "kok",
			Nickname: "max",
			Age:      21,
		},
		CreatedAt: time.Time{},
		VisitedAt: time.Time{},
	}
	return r
}

func (r *userRepository) Create(ctx context.Context, uuid entities.UUID, info entities.UserInfo) error {
	return r.create(ctx, uuid, info)
}

func (r *userRepository) Get(ctx context.Context, uuid entities.UUID) (entities.User, error) {
	return r.get(ctx, uuid)
}

func (r *userRepository) get(ctx context.Context, uuid entities.UUID) (entities.User, error) {
	user, ok := r.userTable[uuid]
	if !ok {
		return entities.User{}, errors.New("db cant get user")
	}
	return user, nil
}

func (r *userRepository) create(ctx context.Context, uuid entities.UUID, info entities.UserInfo) error {
	var (
		user entities.User
	)

	uuid = entities.UUID(fmt.Sprintf("%d", len(r.userTable)))

	user = entities.User{
		UUID:      "",
		Info:      info,
		CreatedAt: time.Now(),
	}

	r.userTable[uuid] = user

	return nil
}
