package main

import (
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

type Link struct {
	ID   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Link string        `json:"link"`
	Code string        `json:"code"`
}

type Jump struct {
	Link bson.ObjectId
	Time time.Time
}

func LinkHandler(w http.ResponseWriter, r *http.Request) {

	link := r.URL.Query().Get("link")
	if link == "" {
		panic("emty link")
	}

	valid := govalidator.IsURL(link)
	if !valid {
		panic("invalid link")
	}

	var linkObj Link
	linkObj.ID = bson.NewObjectId()
	linkObj.Link = link
	linkObj.Code = randStringBytes(10)

	insertIntoCollection(r, "links", linkObj)

	shortedLink := getShortedLink(linkObj.Code)
	writeJSON(w, map[string]string{
		"shortedLink": shortedLink,
		"linkId":      linkObj.ID.Hex(),
	})
}

func StatHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	linkId := vars["linkId"]
	writeJSON(w, map[string]string{"stat": linkId})
}

func JumpHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	linkCode := vars["linkCode"]

	db := getMongoDBFromContext(r)
	collection := db.C("links")

	var linkObj Link
	err := collection.Find(bson.M{"code": linkCode}).One(&linkObj)
	if err != nil {
		http.NotFound(w, r)
	}

	jump := Jump{
		Link: linkObj.ID,
		Time: time.Now(),
	}
	insertIntoCollection(r, "jumps", jump)

	http.Redirect(w, r, linkObj.Link, 302)
}
