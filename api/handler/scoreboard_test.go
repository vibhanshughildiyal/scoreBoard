package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"net/http"
	"scoreBoard/internal/config"
	"scoreBoard/models"
	"scoreBoard/models/dtos"
	"testing"

	handlerMocks "scoreBoard/api/handler/mocks"
	repositoryMocks "scoreBoard/internal/repository/mocks"
	testhelper "scoreBoard/internal/testing"
)

func TestGetScoreBoard(t *testing.T) {
	methodName := "GetScoreBoard"

	storage := &handlerMocks.Repository{}
	gameDao := &repositoryMocks.GameDAO{}
	storage.On("Game").Return(gameDao)

	game := dtos.Game{}
	gameDao.On("WithGameID", mock.Anything).Return(&game, nil)

	scoreBoard := []*models.ScoreBoard{
		{
			Rank: 1,
		},
	}

	gameDao.On("GetLeaderBoardForAllRegion", mock.Anything, mock.Anything, mock.Anything).Return(scoreBoard, nil)
	gameDao.On("GetLeaderBoardForRegion", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(scoreBoard, nil)
	gameDao.On("GetTimedLeaderBoardForAllRegion", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(scoreBoard, nil)
	gameDao.On("GetTimedLeaderBoardForRegion", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(scoreBoard, nil)

	config := &config.Config{}
	var api = ProvideAPI(config, WithRepository(storage))
	(testhelper.HandlerTestCases{
		{
			Title: "should return result",
			Request: testhelper.HandlerRequest{
				Params: gin.Params{{Key: "gid", Value: ""}},
				//Query:  url.Values{"offset": []string{"0"}, "count": []string{"5"}, "from": []string{"1"}, "to": []string{"1"}, "region": []string{"1"}},
			},
			API:          api,
			MethodName:   methodName,
			ExpectedResp: testhelper.HandlerRespAssert{Status: http.StatusOK, Resp: []*models.ScoreBoard{{Rank: 1}}},
		},
	}).Run(t)
}
