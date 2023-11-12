package sessionmgmt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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

	if ServerTest == nil {
		GetServer(false)
	}

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

func GetDummyUser() (user *User) {
	user = &User{
		ID: "2",
	}
	return
}
