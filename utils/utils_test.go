package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestGetNewUUID(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "generate a new uuid",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := GetNewUUID()
			assert.NotEqual(t, response, nil)
		})
	}
}

func TestStringToUuid(t *testing.T) {
	newUuid := GetNewUUID()
	type fields struct {
		stringUuid string
		uuid       uuid.UUID
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "get uuid from string",
			fields:  fields{stringUuid: newUuid.String(), uuid: newUuid},
			wantErr: false,
		},
		{
			name:    "error from invalid uuid string",
			fields:  fields{stringUuid: "invaliduuid"},
			wantErr: true,
		},
		{
			name:    "error from empty string",
			fields:  fields{stringUuid: ""},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, err := StringToUuid(test.fields.stringUuid)
			if test.wantErr {
				assert.Error(t, err.Error)
			} else {
				assert.Equal(t, response, test.fields.uuid)
			}
		})
	}
}

func TestGetQueryOrDefault(t *testing.T) {
	type fields struct {
		key         string
		searchKey   string
		value       string
		defautValue string
	}

	tests := []struct {
		name     string
		fields   fields
		keyExist bool
	}{
		{
			name:     "query param exist with value",
			fields:   fields{key: "key", searchKey: "key", value: "value", defautValue: "defaultValue"},
			keyExist: true,
		},
		{
			name:     "query param does not exist with value",
			fields:   fields{key: "key", searchKey: "keys", defautValue: "defaultValue"},
			keyExist: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := &http.Request{URL: &url.URL{Host: "abc"}}
			q := req.URL.Query()
			q.Add(test.fields.key, test.fields.value)
			req.URL.RawQuery = q.Encode()
			c := &gin.Context{Request: req}
			response := GetQueryOrDefault(c, test.fields.searchKey, test.fields.defautValue)
			if test.keyExist {
				assert.Equal(t, response, test.fields.value)
				assert.NotEqual(t, response, test.fields.defautValue)
			} else {
				assert.Equal(t, response, test.fields.defautValue)
				assert.NotEqual(t, response, test.fields.value)
			}
		})
	}
}
