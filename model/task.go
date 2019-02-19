package model

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	Task struct {
		ID      bson.ObjectId     `json:"id" bson:"_id,omitempty"`
		UserId      string        `json:"UserId" bson:"UserId"`
		TaskName    string        `json:"TaskName" bson:"TaskName"`
		Description string        `json:"Description" bson:"Description"`
		Status 		int        	  `json:"Status" bson:"Status"`
	}
)
