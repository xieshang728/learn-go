package main

import (
	"fmt"
	"learn-go/tree/retiever/mock"
	"learn-go/tree/retiever/really"
	"time"
)

const url = "https://www.imooc.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface{
	Post(url string, form map[string]string) string
}

func post(p Poster){
	p.Post(url, map[string]string{
		"name" :"xieshang",
		"course" :"math",
	})
}

type RetieverPoster interface {
	Retriever
	Poster
}

func session(s RetieverPoster){
	s.Post(url, map[string]string{
		"contents":"this is a content",
	})
}



func download(retriever Retriever) string{
	return retriever.Get("http://www.imooc.com")
}

func inspect(r Retriever){
	fmt.Printf("%T=======>%v",r,r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents===>"+v.Contents)
	case *really.Retiever:
		fmt.Println("UserAgent====>"+v.UserAgent)

	}
}

func main(){
	var r1,r2 Retriever
	r1 = mock.Retriever{"this is a mock Retriever"}

	r2 = &really.Retiever{
		UserAgent:"json/application",
		TimeOut: time.Minute,
	}

	inspect(r1)
	inspect(r2)

	mockRetiever := r1.(mock.Retriever)
	fmt.Println(mockRetiever.Contents)

	//fmt.Println(download(r2))
	//r2 = really.Retiever{UserAgent:"json/application",TimeOut:time.Minute}
	//fmt.Println()
	//fmt.Printf("%T====>%v",r,r)
	//fmt.Println(download(r))

	fmt.Println("=================>>mock Retiever=================")
	session(&mockRetiever)
}