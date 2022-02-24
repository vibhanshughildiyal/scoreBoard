package repository

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	scoreErr "scoreBoard/internal/errors"
	"scoreBoard/models"
	"scoreBoard/models/dtos"
	"scoreBoard/pkg/amqp"
)

type ScoreDAO interface {
	StartConsumingScoreFromAmqp()
}

type scoreDao struct {
	database *gorm.DB
	amqp     amqp.Amqp
}

func (sc scoreDao) StartConsumingScoreFromAmqp() {
	res := make(chan []byte)
	go sc.amqp.ConsumeMessage(res)
	for {
		select {
		case out := <-res:
			var score models.PublishScoreAmqp
			if err := json.Unmarshal(out, &score); err != nil {
				scoreErr.Error(scoreErr.Code("1.0"), err)
			} else {
				//call method to update data in database in a go routine
				go sc.addScore(score)
				log.Info().Msgf("message from amqp:%+v", score)
			}
		}
	}
}

func (sc scoreDao) addScore(payload models.PublishScoreAmqp) *scoreErr.Errors {
	score := dtos.Score{
		Id:        uuid.New(),
		GameId:    payload.GameId,
		PlayerId:  payload.PlayerID,
		Score:     payload.Score,
		CreatedAt: payload.CreatedAt,
	}
	if insertErr := sc.database.Create(&score).Error; insertErr != nil {
		return scoreErr.Error(scoreErr.Code("1.0"), insertErr)
	}
	return nil
}
