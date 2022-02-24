package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	scoreErr "scoreBoard/internal/errors"
	"scoreBoard/models"
	"time"
)

func (a *API) PublishScore(c *gin.Context) (int, interface{}, error) {
	gameId := c.Param("gameID")
	playerId := c.Param("playerID")
	scoreBody := models.PublishScoreBody{}
	if err := c.ShouldBindBodyWith(&scoreBody, binding.JSON); err != nil {
		return http.StatusBadRequest, nil, scoreErr.Error(scoreErr.Code("1.0"), err).Error
	}
	publishScoreAmqp := models.PublishScoreAmqp{
		GameId:    gameId,
		PlayerID:  playerId,
		Score:     scoreBody.Score,
		CreatedAt: time.Now().UTC().Unix(),
	}

	bytes, err := json.Marshal(publishScoreAmqp)
	if err != nil {
		return http.StatusInternalServerError, nil, scoreErr.Error(scoreErr.Code("1.0"), err).Error
	}
	if err := a.Amqp.PublishMessage(bytes); err != nil {
		return http.StatusInternalServerError, nil, err.Error
	}
	return http.StatusCreated, nil, nil
}
