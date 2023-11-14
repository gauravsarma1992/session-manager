package sessionmgmt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gauravsarma1992/src/sessionmgmt"
)

func TestServerStart(t *testing.T) {
	resp, respBody, _ := MakeRequest("GET", "ping", nil)
	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, respBody["message"], "pong")
}

func TestServerGetSessionFailure(t *testing.T) {
	resp, respBody, _ := MakeRequest("GET", "api/sessions/1", nil)
	assert.Equal(t, resp.StatusCode, 401)
	assert.Equal(t, respBody["message"], "No Session found")
	assert.Equal(t, respBody["session"], nil)
}

func TestServerGetSessionSuccess(t *testing.T) {
	server, _ := GetServer(false)
	// Create the session
	session, err := server.sessionStore.AddSession(sessionmgmt.TokenT("1"), GetDummyUser())
	assert.Equal(t, err, nil)

	// Check if session has been created
	resp, respBody, _ := MakeRequest(
		"GET",
		fmt.Sprintf("api/sessions/%s", session.Token),
		nil,
	)
	assert.Equal(t, resp.StatusCode, 200)
	assert.NotEqual(t, respBody["session"], nil)
}

func TestServerRemoveSession(t *testing.T) {
	server, _ := GetServer(false)
	// Create the session
	session, err := server.sessionStore.AddSession(sessionmgmt.TokenT("2"), GetDummyUser())
	assert.Equal(t, err, nil)

	// Delete the session
	resp, respBody, _ := MakeRequest(
		"DELETE",
		fmt.Sprintf("api/sessions/%s", session.Token),
		nil,
	)
	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, respBody["message"], "success")

	// Check if the session got deleted
	resp, respBody, _ = MakeRequest(
		"GET",
		fmt.Sprintf("api/sessions/%s", session.Token),
		nil,
	)
	assert.Equal(t, resp.StatusCode, 401)
}
