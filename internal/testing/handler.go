package testing

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

// HandlerRespAssert set exted values in this form to easily assert against handler response
type HandlerRespAssert struct {
	Status     int
	Resp       interface{}
	HasError   bool
	ErrorCause error
}

// HandlerRequest define handler request to create context out of it
type HandlerRequest struct {
	Body    interface{}
	Headers map[string]string
	Params  gin.Params
	Query   url.Values
}

// HandlerTestCase define handler test cases
type HandlerTestCase struct {
	Title        string
	Request      HandlerRequest
	API          interface{}
	MethodName   string
	ExpectedResp HandlerRespAssert
}

// HandlerTestCases group of HandlerTestCase to run together
type HandlerTestCases []HandlerTestCase

// HasErrorResp usefull assert object when error is expected
var HasErrorResp = &HandlerRespAssert{http.StatusInternalServerError, nil, true, nil}

func ErrorWithCause(err error) *HandlerRespAssert {
	return &HandlerRespAssert{http.StatusInternalServerError, nil, true, err}
}

// Assert handler response against expected response
func (expected HandlerRespAssert) Assert(t *testing.T, status int, resp interface{}, err error) {
	assert.Equal(t, expected.Status, status)
	assert.Equal(t, expected.Resp, resp)
	if expected.HasError {
		assert.Error(t, err)
		if expected.ErrorCause != nil {
			assert.Equal(t, expected.ErrorCause, errors.Cause(err))
		}
	} else {
		assert.NoError(t, err)
	}
}

// GinTestContext generate gin test context with HandlerRequest pre set in the context
func (req *HandlerRequest) GinTestContext() (*gin.Context, *httptest.ResponseRecorder) {
	c, w := GetTestContext()
	if req.Body != nil {
		ginContextAddBody(c, req.Body)
	}

	if req.Headers != nil {
		ginContextAddHeaders(c, req.Headers)
	}

	if req.Params != nil {
		c.Params = req.Params
	}
	if req.Query != nil {
		ginContextAddQueries(c, req.Query)
	}
	return c, w
}

// Run all test cases
func (tCases HandlerTestCases) Run(t *testing.T) {
	for _, tCase := range tCases {
		t.Run(tCase.Title, tCase.Run)
	}
}

// Run single test case
func (tCase HandlerTestCase) Run(t *testing.T) {
	c, _ := tCase.Request.GinTestContext()
	outp := reflect.ValueOf(tCase.API).MethodByName(tCase.MethodName).Call([]reflect.Value{reflect.ValueOf(c)})
	status, resp, errIface := outp[0].Interface(), outp[1].Interface(), outp[2].Interface()
	var err error
	if errIface != nil {
		err = errIface.(error)
	}
	tCase.ExpectedResp.Assert(t, status.(int), resp, err)
}
