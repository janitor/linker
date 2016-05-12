package main

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHandlers(t *testing.T) {
	server := httptest.NewServer(getHandler())
	defer server.Close()
	config.AppHost = server.URL
	initMongo()
	assert := assert.New(t)

	targetUrl := "https://google.com?foo=bar"

	//1. get shorted link
	resp, err := http.Get(config.AppHost + "/link?link=" + url.QueryEscape(targetUrl))
	if err != nil {
		assert.FailNow(err.Error())
	}
	defer resp.Body.Close()

	assert.Equal(resp.StatusCode, 200)

	//2. check response
	decoder := json.NewDecoder(resp.Body)
	var linkResponse LinkResponse
	err = decoder.Decode(&linkResponse)

	if err != nil {
		assert.FailNow(err.Error())
	}

	//3. check redirect
	resp, err = getClientWithoutRedirects().Get(linkResponse.ShortedLink)
	assert.Equal(resp.StatusCode, 302)
	assert.Equal(resp.Header.Get("location"), targetUrl)
}

func getClientWithoutRedirects() *http.Client {
	client := http.Client{
		CheckRedirect: doNotFollowingRedirects,
	}
	return &client
}

func doNotFollowingRedirects(req *http.Request, via []*http.Request) error {
	return errors.New("Redirect")
}
