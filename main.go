package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var config *Configuration

func init() {
	parseOptions()
}

func main() {
	fmt.Println("Ok. Starting...")

	initMongo()
	defer closeMongo()

	http.Handle("/", getHandler())
	http.ListenAndServe(":8000", nil)
}

func getHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/link", LinkHandler)
	router.HandleFunc("/stat/{linkId}", StatHandler)
	router.HandleFunc("/j/{linkCode}", JumpHandler)
	return mongoSessionMiddleware(router)
}

type Configuration struct {
	MongoHost string
	AppHost   string
}

func parseOptions() {
	mongoHost := flag.String("mongo-host", "localhost", "Mongo Host")
	appHost := flag.String("app-host", "http://localhost:8000", "App host")

	flag.Parse()

	conf := Configuration{
		MongoHost: *mongoHost,
		AppHost:   *appHost,
	}
	config = &conf
}
