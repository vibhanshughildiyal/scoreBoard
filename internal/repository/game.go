package repository

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	scoreErr "scoreBoard/internal/errors"
	"scoreBoard/models"
	"scoreBoard/models/dtos"
	"scoreBoard/utils"
)

type GameDAO interface {
	WithGameID(gameId string) (*dtos.Game, *scoreErr.Errors)
	GetLeaderBoardForAllRegion(gameId string, offset, count int) ([]*models.ScoreBoard, *scoreErr.Errors)
	GetLeaderBoardForRegion(gameId, region string, offset, count int) ([]*models.ScoreBoard, *scoreErr.Errors)
	GetTimedLeaderBoardForAllRegion(gameId string, offset, count int, fromTime, toTime string) ([]*models.ScoreBoard, *scoreErr.Errors)
	GetTimedLeaderBoardForRegion(gameId, region string, offset, count int, fromTime, toTime string) ([]*models.ScoreBoard, *scoreErr.Errors)
}

type gameDao struct {
	database *gorm.DB
}

func (g gameDao) WithGameID(gameId string) (*dtos.Game, *scoreErr.Errors) {
	var res dtos.Game
	gameUUID, parseErr := utils.StringToUuid(gameId)
	if parseErr != nil {
		return nil, parseErr
	}
	if err := g.database.Model(&dtos.Game{}).Where(&dtos.Game{Id: gameUUID}).First(&res).Error; err != nil {
		return nil, scoreErr.Error(scoreErr.Code("1.0"), err)
	}
	log.Info().Msgf("%+v", res)
	return &res, nil
}

func (g gameDao) GetLeaderBoardForAllRegion(gameId string, offset, count int) (res []*models.ScoreBoard, boardErr *scoreErr.Errors) {
	err := g.database.Raw(`
					SELECT RANK () OVER (ORDER BY tempRank asc) rank,player_name,player_id,score,region FROM(
					SELECT * FROM(
						SELECT DISTINCT ON (player_id) tempRank,player_name,player_id,score,region FROM (
							SELECT RANK () OVER (ORDER BY score desc, s.created_at ASC) tempRank,p.name AS player_name,s.player_id AS player_id,score,region FROM scores s
								INNER JOIN players p ON (s.player_id=p.id)
								WHERE game_id=$1
						)p
					)q
					ORDER BY tempRank ASC
					)r
					OFFSET $2 
					LIMIT $3;`, gameId, offset, count).Find(&res).Error
	if err != nil {
		return nil, scoreErr.Error(scoreErr.Code("1.0"), err)
	}
	return res, nil
}

func (g gameDao) GetLeaderBoardForRegion(gameId, region string, offset, count int) ([]*models.ScoreBoard, *scoreErr.Errors) {
	var res []*models.ScoreBoard
	err := g.database.Raw(`
					SELECT RANK () OVER (ORDER BY tempRank asc) rank,player_name,player_id,score FROM(
					SELECT * FROM(
						SELECT DISTINCT ON (player_id) tempRank,player_name,player_id,score FROM (
							SELECT RANK () OVER (ORDER BY score desc, s.created_at ASC) tempRank,p.name AS player_name,s.player_id AS player_id,score FROM scores s
								INNER JOIN players p ON (s.player_id=p.id)
								WHERE game_id=$1 AND s.region=$2
						)p
					)q
					ORDER BY tempRank ASC
					)r
					OFFSET $3 
					LIMIT $4;`, gameId, region, offset, count).Find(&res).Error
	if err != nil {
		return nil, scoreErr.Error(scoreErr.Code("1.0"), err)
	}
	return res, nil
}

func (g gameDao) GetTimedLeaderBoardForAllRegion(gameId string, offset, count int, fromTime, toTime string) ([]*models.ScoreBoard, *scoreErr.Errors) {
	var res []*models.ScoreBoard
	err := g.database.Raw(`
					SELECT RANK () OVER (ORDER BY tempRank asc) rank,player_name,player_id,score,region FROM(
					SELECT * FROM(
						SELECT DISTINCT ON (player_id) tempRank,player_name,player_id,score,region FROM (
							SELECT RANK () OVER (ORDER BY score desc, s.created_at ASC) tempRank,p.name AS player_name,s.player_id AS player_id,score,region FROM scores s
								INNER JOIN players p ON (s.player_id=p.id)
								WHERE game_id=$1 AND s.created_at BETWEEN $4 AND $5
						)p
					)q
					ORDER BY tempRank ASC
					)r
					OFFSET $2 
					LIMIT $3;`, gameId, offset, count, fromTime, toTime).Find(&res).Error
	if err != nil {
		return nil, scoreErr.Error(scoreErr.Code("1.0"), err)
	}
	return res, nil
}

func (g gameDao) GetTimedLeaderBoardForRegion(gameId, region string, offset, count int, fromTime, toTime string) ([]*models.ScoreBoard, *scoreErr.Errors) {
	var res []*models.ScoreBoard
	err := g.database.Raw(`
					SELECT RANK () OVER (ORDER BY tempRank asc) rank,player_name,player_id,score FROM(
					SELECT * FROM(
						SELECT DISTINCT ON (player_id) tempRank,player_name,player_id,score FROM (
							SELECT RANK () OVER (ORDER BY score desc, s.created_at ASC) tempRank,p.name AS player_name,s.player_id AS player_id,score FROM scores s
								INNER JOIN players p ON (s.player_id=p.id)
								WHERE game_id=$1 AND s.region=$2 AND s.created_at BETWEEN $5 AND $6
						)p
					)q
					ORDER BY tempRank ASC
					)r
					OFFSET $3 
					LIMIT $4;`, gameId, region, offset, count, fromTime, toTime).Find(&res).Error
	if err != nil {
		return nil, scoreErr.Error(scoreErr.Code("1.0"), err)
	}
	return res, nil
}
