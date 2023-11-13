package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var httpClient = &http.Client{Timeout: 5 * time.Second}
var limiter = make(chan struct{}, 200)

type Client struct {
}

type ResponseObjectStatus struct {
	Id     int  `json:"id"`
	Status bool `json:"online"`
}

func (c *Client) getStatus(id int) bool {
	select {
	case limiter <- struct{}{}:
		break
	case <-time.After(30 * time.Second):
		return false
	}
	defer func() {
		<-limiter
	}()
	requestUrl := fmt.Sprintf("http://localhost:9010/objects/%d", id)
	resp, err := httpClient.Get(requestUrl)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var responseObjectStatus ResponseObjectStatus
	err = json.NewDecoder(resp.Body).Decode(&responseObjectStatus)
	if err != nil {
		fmt.Println(err)
	}
	return responseObjectStatus.Status
}
