package main

import (
	"encoding/json"
	"fmt"
)


type Person struct {
	Age int `json:"age"`
	Name string `json:"name"`
	Address string `json:"address"`
}

func main(){
	encode()
	decode()
}

func encode(){
	fmt.Println("------------encode------------")
	p := Person{Age:24,Name:"xieshang",Address:"hubei"}
	json,_ := json.Marshal(p)
	fmt.Printf(string(json))
	fmt.Println()
}


func decode(){
	fmt.Println("-----------decode------------")
	jsonStr := `{"age":21,"name":"xieshang","address":"hubei"}`
	p := Person{}

	//序列化时用引用传递
	json.Unmarshal([]byte(jsonStr),&p)
	fmt.Println(p)
}