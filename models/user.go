package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id    bson.ObjectId `json:"id" bson:"_id"`
	Score int           `json:"score" bson:"score"`
	Name  string        `josn:"name" bson:"name"`
	Age   int           `json:"age" bson:"age"`
}
