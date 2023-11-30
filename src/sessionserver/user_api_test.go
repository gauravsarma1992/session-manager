package sessionserver

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	sessionmgmt "github.com/gauravsarma1992/src/sessionmgmt"
)

func testServerGetUser(user *sessionmgmt.User, t *testing.T) {
	resp, respBody, err := MakeRequest(
		"GET",
		fmt.Sprintf("api/users/%d", user.ID),
		nil,
	)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, uint(respBody["user"].(map[string]interface{})["id"].(float64)), user.ID)
}

func testServerAddUser(user *sessionmgmt.User, t *testing.T) {
	var (
		err error
	)
	// Create the user
	resp, respBody, err := MakeRequest(
		"POST",
		"api/users",
		user,
	)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, respBody["message"], "success")
	user.ID = uint(respBody["user"].(map[string]interface{})["id"].(float64))
}

func testServerUpdateUser(user *sessionmgmt.User, t *testing.T) {
	var (
		err error
	)
	// Create the user
	user.Mobile = "84343545555"
	resp, respBody, err := MakeRequest(
		"PUT",
		fmt.Sprintf("api/users/%d", user.ID),
		user,
	)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, respBody["message"], "success")
	assert.Equal(t, user.Mobile, respBody["user"].(map[string]interface{})["mobile"].(string))

}

func testServerDeleteUser(user *sessionmgmt.User, t *testing.T) {
	var (
		err error
	)
	// Delete the user
	resp, respBody, err := MakeRequest(
		"DELETE",
		fmt.Sprintf("api/users/%d", user.ID),
		user,
	)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, respBody["message"], "success")

}

func testServerLoginUserSuccess(user *sessionmgmt.User, t *testing.T) {
	var (
		err error
	)
	loginRequest := &LoginRequest{
		Username: user.Username,
		Password: user.Password,
	}
	// Create the user
	resp, respBody, err := MakeRequest(
		"POST",
		"api/users/login",
		loginRequest,
	)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, respBody["message"], "success")

}

func testServerLoginUserFailure(user *sessionmgmt.User, t *testing.T) {
	var (
		err error
	)
	loginRequest := &LoginRequest{
		Username: user.Username,
		Password: "dummy",
	}
	// Create the user
	resp, respBody, err := MakeRequest(
		"POST",
		"api/users/login",
		loginRequest,
	)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, respBody["message"], "success")

}

func TestUserFlow(t *testing.T) {
	user := &sessionmgmt.User{
		Username: "goku",
		Password: "password",
	}
	server, _ := GetServer(false)
	server.Db.Exec("DELETE FROM users")

	testServerAddUser(user, t)
	testServerGetUser(user, t)
	testServerLoginUserSuccess(user, t)
	testServerLoginUserFailure(user, t)
	testServerUpdateUser(user, t)
	testServerDeleteUser(user, t)
}
