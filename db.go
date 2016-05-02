package main

import (
	"github.com/gorilla/context"
	"gopkg.in/mgo.v2"
	"net/http"
)

var mongoSession *mgo.Session

func mongoSessionMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := getMongoSession()
		defer session.Close()
		context.Set(r, CONTEXT_KEY_MONGO_SESSION, session)
		context.Set(r, CONTEXT_KEY_MONGO_DB, session.DB(config.MongoDB))
		h.ServeHTTP(w, r)
	})
}

func insertIntoCollection(r *http.Request, collectionName string, obj interface{}) {
	db := getMongoDBFromContext(r)
	collection := db.C(collectionName)
	collection.Insert(obj)
}

func getMongoDBFromContext(r *http.Request) *mgo.Database {
	rv := context.Get(r, CONTEXT_KEY_MONGO_DB)
	return rv.(*mgo.Database)
}

func getMongoSession() *mgo.Session {
	return mongoSession.Clone()
}
