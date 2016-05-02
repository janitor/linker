package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"gopkg.in/mgo.v2"
)

const MONGO_DB = "linker"

const CONTEXT_KEY_MONGO_SESSION = "mongo_session"
const CONTEXT_KEY_MONGO_DB = "mongo_db"

func main() {
	session, err := mgo.Dial("localhost")
 	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	mongoSession = session

	router := mux.NewRouter()
	router.HandleFunc("/link", LinkHandler)
	router.HandleFunc("/stat/{linkId}", StatHandler)

	http.Handle("/", mongoSessionMiddleware(router))
	http.ListenAndServe(":8000", nil)
}

