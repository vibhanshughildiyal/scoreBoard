package testing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
)

// GetTestContext returns a gin.Context using gin.CreateTestContext(..) which
// in turn returns a gin.Context to be used for testing purpose
func GetTestContext() (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
	return c, w
}

// GetTestContextWithBody returns a gin.Context using gin.CreateTestContext(..)
// along with a pre bound request body
func GetTestContextWithBody(jsonObj interface{}) (*gin.Context, *httptest.ResponseRecorder) {
	c, w := GetTestContext()
	ginContextAddBody(c, jsonObj)
	return c, w
}

func ginContextAddBody(c *gin.Context, jsonObj interface{}) {
	if jsonObj != nil {
		if jsonValue, err := json.Marshal(jsonObj); err == nil {
			c.Request, _ = http.NewRequest("", "", bytes.NewBuffer(jsonValue))
		}
	}
}

func ginContextAddHeaders(c *gin.Context, headers map[string]string) {
	for k, v := range headers {
		c.Request.Header.Add(k, v)
	}
}

func ginContextAddQueries(c *gin.Context, vals url.Values) {
	c.Request.URL = &url.URL{RawQuery: vals.Encode()}
}
