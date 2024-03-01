package api

import (
	"basic-backend/internal/entities"
)

type User struct {
	UUID entities.UUID `json:"uuid"`
	Info UserInfo      `json:"info"`
}

type UserInfo struct {
	Login    string `json:"login"`
	Nickname string `json:"nickname"`
	Age      int    `json:"age"`
}
