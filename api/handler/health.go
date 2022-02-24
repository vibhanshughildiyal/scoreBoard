package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *API) HealthCheck(c *gin.Context) (int, interface{}, error) {
	return http.StatusOK, nil, nil
}
