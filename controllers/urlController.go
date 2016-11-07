package controllers

import (
	"encoding/json"
	"net/http"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
	"github.com/vineetdaniel/AiOps/apiv1/common"
	"github.com/vineetdaniel/AiOps/apiv1/data"
)

//Handler for HTTP Post - "urls"
//Insert a new url

func CreateUrl(w http.ResponseWriter, r *http.Request) {
	var dataResource UrlResource
	//Decode the incoming url json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w,
			err,
			"Invalid Task Data",
			500,
		)
		return

	}
	url := &dataResource.Data
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("urls")
	repo := &data.UrlRepository{c}
	//Insert url document
	repo.Create(url)
	if j, err := json.Marshal(UrlResource{Data: *url}); err != nil {
		common.DisplayAppError(w,
			err,
			"An expected error has occured",
			500,
		)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}

//Handler for HTTP Get - "/urls"
//Returns all URL documents

func GetUrls(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("urls")
	repo := &data.UrlRepository{c}
	urls := repo.GetAll()
	j, err := json.Marshal(UrlsResource{Data: urls})
	if err != nil {
		common.DisplayAppError(w,
			err,
			"An expected error has occured",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

//Handler for HTTP Get - "/urls/users/{id}"
//Return all urls created by user

func GetUrlsByUser(w http.ResponseWriter, r *http.Request) {
	//get id from an incoming url
	vars := mux.Vars(r)
	user := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("urls")
	repo := &data.UrlRepository{c}
	urls := repo.GetByUser(user)
	j, err := json.Marshal(UrlsResource{Data: urls})
	if err != nil {
		common.DisplayAppError(w,
			err,
			"An expected error has occured",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
