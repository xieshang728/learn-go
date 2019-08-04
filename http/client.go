package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main(){
	resp,err := http.Get("http://www.imooc.com")
	if err != nil {
		log.Fatalf("get imooc error %s\n",err)
	}

	defer resp.Body.Close()

	s,err := httputil.DumpResponse(resp,true)
	if err != nil{
		log.Fatalf("parse body err %s\n",err)
	}
	fmt.Printf("response %s",s)
}
