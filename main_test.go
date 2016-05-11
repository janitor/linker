package main

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestHandlers(t *testing.T) {
	assert := assert.New(t)

	targetUrl := "https://google.com"
	targetUrlEscaped := url.QueryEscape(targetUrl)

	resp, err := http.Get("http://localhost:8000/link?link=" + targetUrlEscaped)
	if err != nil {
		assert.Error(err)
	}
	defer resp.Body.Close()

	assert.Equal(resp.StatusCode, 200)

	decoder := json.NewDecoder(resp.Body)
	var linkResponse LinkResponse
	err = decoder.Decode(&linkResponse)

	if err != nil {
		assert.Error(err)
	}

	resp, err = http.Get(linkResponse.ShortedLink)
	if err != nil {
		assert.Error(err)
	}

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return errors.New("Redirect")
		},
	}

	resp, err = client.Get(linkResponse.ShortedLink)
	if err != nil {
		if resp.StatusCode == 302 {
			assert.Equal(resp.Header.Get("location"), targetUrl)
		} else {
			assert.Error(err)
		}
	} else {
		assert.Error(errors.New("Where is my redirect ?"))
	}
}
