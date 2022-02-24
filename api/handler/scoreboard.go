package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	scoreErr "scoreBoard/internal/errors"
	"scoreBoard/models"
	"scoreBoard/utils"
	"strconv"
	"time"
)

const (
	defaultLeaderboardOffset = 0
	defaultLeaderboardCount  = 5
	dateTimeFormat           = "2006-01-02"
)

func (a *API) GetScoreBoard(c *gin.Context) (int, interface{}, error) {
	gameId := c.Param("gid")
	offsetQuery := utils.GetQueryOrDefault(c, "offset", "0")
	countQuery := utils.GetQueryOrDefault(c, "count", "5")
	fromTimeQuery := utils.GetQueryOrDefault(c, "from", "")
	toTimeQuery := utils.GetQueryOrDefault(c, "to", "")
	region := utils.GetQueryOrDefault(c, "region", "")

	//get game details
	game, gameErr := a.Storage.Game().WithGameID(gameId)
	if gameErr != nil {
		return http.StatusInternalServerError, nil, gameErr.Error
	}
	//check if game exist or not
	if game == nil {
		return http.StatusBadRequest, nil, scoreErr.Error(scoreErr.Code("3.0")).Error
	}

	//convert if offset to int; else assign default offset value
	offset, strConvErr := strconv.Atoi(offsetQuery)
	if strConvErr != nil {
		offset = defaultLeaderboardOffset
	}

	//convert if count to int; else assign default count value
	count, strConvErr := strconv.Atoi(countQuery)
	if strConvErr != nil {
		count = defaultLeaderboardCount
	}

	var fromTime, toTime time.Time
	var err error
	if fromTimeQuery != "" {
		fmt.Println("fromTimeQuery, toTimeQuery)")
		fmt.Println(fromTimeQuery, toTimeQuery)
		fromTime, err = time.Parse(dateTimeFormat, fromTimeQuery)
		if err != nil {
			log.Error().Msgf("unable to read from date from query: %s", fromTimeQuery)
			fromTimeQuery = ""
		}
	}
	if toTimeQuery != "" {
		fmt.Println("fromTimeQuery, toTimeQuery)")
		fmt.Println(fromTimeQuery, toTimeQuery)
		toTime, err = time.Parse(dateTimeFormat, toTimeQuery)
		if err != nil {
			log.Error().Msgf("unable to read to date from query: %s", toTimeQuery)
			toTimeQuery = ""
		}
	}

	var leaderBoard []*models.ScoreBoard
	var leaderBoardErr *scoreErr.Errors

	//get scoreboard for all time
	if fromTimeQuery == "" || toTimeQuery == "" {
		if region == "" {
			leaderBoard, leaderBoardErr = a.Storage.Game().GetLeaderBoardForAllRegion(gameId, offset, count)
		} else {
			leaderBoard, leaderBoardErr = a.Storage.Game().GetLeaderBoardForRegion(gameId, region, offset, count)
		}
	} else {
		//swap if toTime comes before fromTime
		if toTime.Before(fromTime) {
			fromTime, toTime = toTime, fromTime
		}
		fromTimeQuery = fmt.Sprintf("%s 00:00:00", fromTime.UTC().Format(dateTimeFormat))
		toTimeQuery = fmt.Sprintf("%s 23:59:59", toTime.UTC().Format(dateTimeFormat))
		//get scoreboard for a time frame
		if region == "" {
			leaderBoard, leaderBoardErr = a.Storage.Game().GetTimedLeaderBoardForAllRegion(gameId, offset, count, fromTimeQuery, toTimeQuery)
		} else {
			leaderBoard, leaderBoardErr = a.Storage.Game().GetTimedLeaderBoardForRegion(gameId, region, offset, count, fromTimeQuery, toTimeQuery)
		}
	}
	if leaderBoardErr != nil {
		log.Error().Err(leaderBoardErr.Error)
		return http.StatusInternalServerError, nil, leaderBoardErr.Error
	}

	return http.StatusOK, leaderBoard, nil
}
