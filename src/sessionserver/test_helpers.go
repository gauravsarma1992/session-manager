package sessionserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gauravsarma1992/src/sessionmgmt"
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
			ServerTest.config.Server.Host,
			ServerTest.config.Server.Port,
		)

		go ServerTest.Run()
		time.Sleep(1 * time.Second)
	}
	server = ServerTest
	return
}

func MakeRequest(method, api string, payload interface{}) (resp *http.Response, respBody map[string]interface{}, err error) {

	var (
		payloadB []byte
	)

	if ServerTest == nil {
		GetServer(false)
	}

	if payload != nil {
		if payloadB, err = json.Marshal(payload); err != nil {
			return
		}

	}

	req := &http.Request{}
	client := &http.Client{}
	url := fmt.Sprintf("%s/%s", BaseUrl, api)

	if req, err = http.NewRequest(method, url, bytes.NewBuffer(payloadB)); err != nil {
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

func GetDummyUser() (user *sessionmgmt.User) {
	user = &sessionmgmt.User{
		ID: "2",
	}
	return
}
