package sessionmgmt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	ServerTest *Server
	BaseUrl    string
)

func GetServer(forceNew bool) (server *Server, err error) {
	if forceNew == true || ServerTest == nil {
		ServerTest, _ = NewServer(nil)
		BaseUrl = fmt.Sprintf(
			"http://%s:%s",
			ServerTest.config.Host,
			ServerTest.config.Port,
		)

		go ServerTest.Run()
		time.Sleep(1 * time.Second)
	}
	server = ServerTest
	return
}

func MakeRequest(method, api string, payload interface{}) (resp *http.Response, respBody map[string]interface{}, err error) {
	req := &http.Request{}
	client := &http.Client{}
	url := fmt.Sprintf("%s/%s", BaseUrl, api)

	if req, err = http.NewRequest(method, url, nil); err != nil {
		return
	}
	if resp, err = client.Do(req); err != nil {
		return
	}
	respBody = make(map[string]interface{})

	respB, _ := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(respB, &respBody); err != nil {
		return
	}

	return
}

func TestServerStart(t *testing.T) {
	_, err := GetServer(true)
	assert.Equal(t, err, nil)

	_, respBody, err := MakeRequest("GET", "ping", nil)
	assert.Equal(t, err, nil)
	assert.Equal(t, respBody["message"], "pong")
}

func TestServerGetSessionNotPresent(t *testing.T) {
	_, err := GetServer(false)
	assert.Equal(t, err, nil)

	_, respBody, err := MakeRequest("GET", "api/sessions/1", nil)
	assert.Equal(t, err, nil)
	assert.Equal(t, respBody["message"], "success")
}
