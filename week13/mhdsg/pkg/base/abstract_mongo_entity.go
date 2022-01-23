package base

import "time"

type AbstractMongoEntity struct {
	Id string `json:"_id"`
}

type AbstractMongoEntendEntity struct {
	AbstractMongoEntity
	Createtime time.Time `json:"createtime"`
	Updatetime time.Time `json:"updatetime"`
}
