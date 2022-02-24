package models

import "github.com/google/uuid"

type ScoreBoard struct {
	Rank       int
	PlayerId   uuid.UUID
	PlayerName string
	Score      int
	Region     string `json:"region,omitempty"`
}
