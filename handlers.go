package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"github.com/asaskevich/govalidator"
)

type Link struct {
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Link string `json:"link"`
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

	insertIntoCollection(r, "links", linkObj)
	writeJSON(w, map[string]string{"link": link, "linkId": linkObj.ID.Hex()})
}

func StatHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	linkId := vars["linkId"]
	writeJSON(w, map[string]string{"stat": linkId})
}

func writeJSON(w http.ResponseWriter, data interface{}) error {
	jsn, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsn)
	return nil
}
