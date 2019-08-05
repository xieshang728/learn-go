package main

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"io/ioutil"
	"net/http"
)

func main(){
	resp,err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil{
		log.Error("Crawler error %s",err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error("Crawler error %s",resp.StatusCode)
		return
	}

	result,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Error("Crawler error %s",err)
	}

	fmt.Printf("%s",result)

}
