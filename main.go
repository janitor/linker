package main

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"net/http"
)

const CONTEXT_KEY_MONGO_SESSION = "mongo_session"
const CONTEXT_KEY_MONGO_DB = "mongo_db"

func main() {
	loadConfig()

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

	router.HandleFunc("/g/{linkCode}", GotoHandler)

	http.Handle("/", mongoSessionMiddleware(router))
	http.ListenAndServe(":8000", nil)
}

func getShortedLink(linkCode string) string {
	return config.AppProtocol + "://" + config.AppHost + "/g/" + linkCode
}
