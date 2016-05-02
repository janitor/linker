package main

import (
	"encoding/json"
	"os"
)

var config *Configuration

type Configuration struct {
	MongoHost string
	MongoDB   string

	AppProtocol string
	AppHost     string
}

func loadConfig() {
	configFile, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(configFile)
	var conf Configuration
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err)
	}
	config = &conf
}
