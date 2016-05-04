package main

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"net/http"
	"flag"
)

const CONTEXT_KEY_MONGO_SESSION = "mongo_session"
const CONTEXT_KEY_MONGO_DB = "mongo_db"

var config *Configuration

func init() {
	parseOptions()
}

func main() {
	session, err := mgo.Dial(config.MongoHost)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	mongoSession = session

	router := mux.NewRouter()
	router.HandleFunc("/link", LinkHandler)
	router.HandleFunc("/stat/{linkId}", StatHandler)

	router.HandleFunc("/j/{linkCode}", JumpHandler)

	http.Handle("/", mongoSessionMiddleware(router))
	http.ListenAndServe(config.AppHost, nil)
}

func getShortedLink(linkCode string) string {
	return "http://" + config.AppHost + "/j/" + linkCode
}

type Configuration struct {
	MongoHost string
	AppHost     string
}

func parseOptions () {
	mongoHost := flag.String("mongo-host", "localhost", "Mongo Host")
	appHost := flag.String("app-host", "127.0.0.1:8000", "App host")

	flag.Parse()

	conf := Configuration{
		MongoHost: *mongoHost,
		AppHost: *appHost,
	}
	config = &conf
}
