package redisutil

import (
	"encoding/json"
	"fmt"
	"learn-go/model"
	"testing"
)



func TestRedisGet(t *testing.T){
	result := Get("username")
	fmt.Printf("result %q",result)
}

func TestRedisSet(t *testing.T){
	//user := &model.User{Username:"xieshang",Age:21,Address:"beijing",Birthday:"1995-07-28"}
	//body,err := json.Marshal(user)
	//if err != nil{
	//	log.Error("parse json error %s",err.Error())
	//}
	//y
	//Set("user",string(body))

	var user model.User
	body := Get("user")
	fmt.Println(body)
	json.Unmarshal([]byte(body),&user)
	fmt.Println(user)
}


func TestDel(t *testing.T){
	Del("username")
}

func TestSet(t *testing.T) {
	SetForExpire("age","22",3000)
}

func TestSetnx(t *testing.T){
	result := SetNX("birthday","1995-07-28")
	fmt.Println(result)
}