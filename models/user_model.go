package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id	   	 primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   	 string             `json:"name,omitempty" bson:"name,omitempty"`
	Location string       		`json:"location,omitempty" bson:"location,omitempty"`
	Title	 string  			`json:"title,omitempty" bson:"title,omitempty"`
	}