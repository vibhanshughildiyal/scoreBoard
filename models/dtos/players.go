package dtos

import (
	"github.com/google/uuid"
	"time"
)

type Player struct {
	Id        uuid.UUID
	Name      string
	Email     string
	Mobile    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
