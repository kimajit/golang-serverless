package ccitrigger

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var statusCode = "200 OK"
 
type mockclient struct{}

func (client mockclient) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		Status:     statusCode,
		Body:       io.NopCloser(strings.NewReader(`{}`)),
		StatusCode: http.StatusOK,
	}, nil
}

func TestPost(t *testing.T) {
	body := strings.NewReader(`{
		"Status" : "test"
	}`)

	t.Run("Test for successfull trigger", func(t *testing.T) {
		status, err := PostRequestToCircleCI("https://test.com", body, "my-token", mockclient{})
		assert.Nil(t, err)
		assert.Equal(t, status, statusCode)
	})
	t.Run("Test for invalid url", func(t *testing.T) {
		status, err := PostRequestToCircleCI("", body, "my-token", mockclient{})
		assert.NotNil(t, status)
		assert.Nil(t, err)
	})
}
