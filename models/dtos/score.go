package dtos

import "github.com/google/uuid"

type Score struct {
	Id        uuid.UUID
	GameId    string
	PlayerId  string
	Score     int64
	Region    string
	CreatedAt int64
}
