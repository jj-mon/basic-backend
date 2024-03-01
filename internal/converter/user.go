package converter

import (
	"basic-backend/internal/api"
	"basic-backend/internal/entities"
)

func ToUserFromService(in *entities.User) *api.User {
	return &api.User{
		UUID: in.UUID,
		Info: api.UserInfo{
			Login:    in.Info.Login,
			Nickname: in.Info.Nickname,
			Age:      in.Info.Age,
		},
	}
}
