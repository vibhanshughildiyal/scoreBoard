package repository

import (
	"gorm.io/gorm"
	"scoreBoard/pkg/amqp"
)

type Storage struct {
	database *gorm.DB
	amqp     amqp.Amqp
}

func (s *Storage) Game() GameDAO {
	return gameDao{database: s.database}
}

func (s *Storage) Player() PlayerDAO {
	return playerDao{database: s.database}
}

func (s *Storage) Score() ScoreDAO {
	return scoreDao{database: s.database, amqp: s.amqp}
}

func ProvideStorage(database *gorm.DB, amqp amqp.Amqp) *Storage {
	return &Storage{
		database: database,
		amqp:     amqp,
	}
}
