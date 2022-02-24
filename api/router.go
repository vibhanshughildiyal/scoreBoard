package api

import (
	"github.com/gin-gonic/gin"
	"scoreBoard/api/handler"
	"scoreBoard/api/middleware"
	scoreErr "scoreBoard/internal/errors"
)

func NewRouter(api *handler.API) *gin.Engine {
	router := gin.New()

	router.Use(middleware.CORS())
	router.GET("/healthCheck", scoreErr.Advice(api.HealthCheck))

	game := router.Group("/game")
	{
		game.POST(":gid/player/:pid", scoreErr.Advice(api.PublishScore))
		game.GET(":gid/scoreBoard", scoreErr.Advice(api.GetScoreBoard))
	}

	return router
}
