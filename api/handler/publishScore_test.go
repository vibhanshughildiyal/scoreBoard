package handler

/*
import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"net/http"
	handlerMocks "scoreBoard/api/handler/mocks"
	"scoreBoard/internal/config"
	repositoryMocks "scoreBoard/internal/repository/mocks"
	testhelper "scoreBoard/internal/testing"
	"scoreBoard/models"
	"scoreBoard/models/dtos"
	amqpMocks "scoreBoard/pkg/amqp/mocks"
	"testing"
)

func TestPublishScore(t *testing.T) {
	methodName := "PublishScore"

	storage := &handlerMocks.Repository{}
	amqpMock := &amqpMocks.Amqp{}
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
	var api = ProvideAPI(config, WithRepository(storage), WithAmqp(amqpMock))
	(testhelper.HandlerTestCases{
		{
			Title: "should return result",
			Request: testhelper.HandlerRequest{
				Params: gin.Params{{Key: "gameID", Value: "123"}, {Key: "playerID", Value: "123"}},
			},
			API:          api,
			MethodName:   methodName,
			ExpectedResp: testhelper.HandlerRespAssert{Status: http.StatusCreated, Resp: []*models.ScoreBoard{{Rank: 1}}},
		},
	}).Run(t)
}
*/
