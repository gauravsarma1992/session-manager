package sessionmgmt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerStart(t *testing.T) {
	resp, respBody, _ := MakeRequest("GET", "ping", nil)
	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, respBody["message"], "pong")
}

func TestServerGetSessionFailure(t *testing.T) {
	resp, respBody, _ := MakeRequest("GET", "api/sessions/1", nil)
	assert.Equal(t, resp.StatusCode, 500)
	assert.Equal(t, respBody["message"], "No Session found")
	assert.Equal(t, respBody["session"], nil)
}

func TestServerGetSessionSuccess(t *testing.T) {
	server, _ := GetServer(false)
	session, err := server.sessionStore.AddSession(TokenT("1"), GetDummyUser())
	assert.Equal(t, err, nil)
	resp, respBody, _ := MakeRequest(
		"GET",
		fmt.Sprintf("api/sessions/%s", session.Token),
		nil,
	)
	assert.Equal(t, resp.StatusCode, 200)
	assert.NotEqual(t, respBody["session"], nil)
}
