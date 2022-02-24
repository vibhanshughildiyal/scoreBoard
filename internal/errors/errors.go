package serviceerrors

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"reflect"
)

type Code string
type Errors struct {
	Code           Code
	Error          error
	Message        string
	AdditionalInfo []interface{}
}

type ClientError struct {
	Code                Code   `json:"code"`
	Description         string `json:"description"`
	DetailedDescription string `json:"detailedDescription"`
}

// Advice is a decorator which controllers can use to delegate handling errors &
// returning responses
func Advice(api func(*gin.Context) (int, interface{}, error)) func(c *gin.Context) {
	return func(c *gin.Context) {
		statusCode, response, err := api(c)
		if err == nil {
			c.JSON(statusCode, response)
			return
		}
		errResponse := getErrCodeDetails(Error(err))
		c.AbortWithStatusJSON(statusCode, errResponse)
	}
}

func getErrCodeDetails(err *Errors) ClientError {
	var res ClientError
	res.Code = (*err).Code
	if errDetail, ok := ErrCodes[(*err).Code]; ok {
		res.Description, res.DetailedDescription = errDetail.Description, errDetail.DetailDescription
	}
	return res
}

func Error(errParams ...interface{}) *Errors {
	err := &Errors{}
	for _, errObj := range errParams {
		addError(err, errObj)
	}
	log.Error().Msgf("%+v", err)
	return err
}

func addError(err *Errors, obj interface{}) {
	switch obj.(type) {
	case Code:
		(*err).Code = obj.(Code)
	case string:
		(*err).Message = fmt.Sprintln((*err).Message, obj.(string))
	case error:
		(*err).Error = obj.(error)
	case *Errors:
		v := reflect.ValueOf(obj.(*Errors))
		x := v.Elem()
		for i := 0; i < x.NumField(); i++ {
			addError(err, x.Field(i).Interface())
		}
		(*err).Code = obj.(Code)
	default:
		(*err).AdditionalInfo = append((*err).AdditionalInfo, obj)
	}
}
