package data

import (
	"time"

	"github.com/vineetdaniel/AiOps/apiv1/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UrlRepository struct {
	C *mgo.Collection
}

func (r *UrlRepository) Create(url *models.Url) error {
	obj_id := bson.NewObjectId()
	url.Id = obj_id
	url.CreatedOn = time.Now()
	url.Status = "Created"
	err := r.C.Insert(&url)
	return err
}

func (r *UrlRepository) Update(url *models.Url) error {
	//partial update on MongoDB
	err := r.C.Update(bson.M{"_id": url.Id},
		bson.M{"$set": bson.M{
			"name":        url.Name,
			"description": url.Description,
			"interval":    url.Interval,
			"status":      url.Status,
			"tags":        url.Tags,
		}})
	return err
}

func (r *UrlRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (r *UrlRepository) GetByUser(user string) []models.Url {
	var urls []models.Url
	iter := r.C.Find(bson.M{"createdby": user}).Iter()
	result := models.Url{}
	for iter.Next(&result) {
		urls = append(urls, result)
	}
	return urls
}

func (r *UrlRepository) GetAll() []models.Url {
	var urls []models.Url
	iter := r.C.Find(nil).Iter()
	result := models.Url{}
	for iter.Next(&result) {
		urls = append(urls, result)
	}
	return urls
}
