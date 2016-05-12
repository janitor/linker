package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

func writeJSON(w http.ResponseWriter, data interface{}) error {
	jsn, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsn)
	return nil
}

func randStringBytes(n int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func getShortedLink(linkCode string) string {
	return config.AppHost + "/j/" + linkCode
}
