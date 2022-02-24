package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	scoreErr "scoreBoard/internal/errors"
)

// StringToUuid returns a pointer to the string value passed in.
func StringToUuid(v string) (uuid.UUID, *scoreErr.Errors) {
	parsedUuid, err := uuid.Parse(v)
	if err != nil {
		log.Info().Msg("in error")
		log.Error().Err(err)
		return parsedUuid, scoreErr.Error(scoreErr.Code("1.0"), err)
	}
	return parsedUuid, nil
}
func GetNewUUID() uuid.UUID {
	return uuid.New()
}
func GetQueryOrDefault(c *gin.Context, key, defaultValue string) string {
	value := c.Query(key)
	if value == "" {
		value = defaultValue
	}
	return value
}
