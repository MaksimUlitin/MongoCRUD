package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	First_name string        `json:"first_name" bson:"first_name"`
	Last_name  string        `json:"last_name" bson:"last_name"`
	Age        int16         `json:"last_name" bson:"last_name"`
}
