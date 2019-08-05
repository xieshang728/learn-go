package entity

import (
	"encoding/json"
	"github.com/gpmgo/gopm/modules/log"
)

type Book struct {
	Id string `json:"id"`
	Author string `json:"author"`
	Title string `json:'title'`
	Year string `json:'year'`
	Publish string `json:'publish'`
}

func (b Book) SetId(id string){
	b.Id = id
}

func (b Book) String() string{
	json,err := json.Marshal(b)
	if err != nil{
		log.Error("parse json err %s ",err)
	}
	return string(json)
}