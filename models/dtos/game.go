package dtos

import "github.com/google/uuid"

type Game struct {
	Id   uuid.UUID
	Name string
}
