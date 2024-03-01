package entities

import "time"

type UUID string

type User struct {
	UUID      UUID
	Info      UserInfo
	CreatedAt time.Time
	VisitedAt time.Time
}

type UserInfo struct {
	Login    string
	Nickname string
	Age      int
}
