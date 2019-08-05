package main

import (
	"bufio"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
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

	e := determineEncoding(resp.Body)

	mineReader := transform.NewReader(resp.Body,e.NewDecoder())

	result,err := ioutil.ReadAll(mineReader)

	if err != nil{
		log.Error("Crawler error %s",err)
	}

	fmt.Printf("%s",result)

}

func determineEncoding(
	r io.Reader) encoding.Encoding{
		bytes, err := bufio.NewReader(r).Peek(1024)
		if err != nil{
			log.Error("Main.determineEncoding.error %s",err)
		}

		e,_,_ := charset.DetermineEncoding(bytes,"")
		log.Info("Main determineEncoding.result %s",e)
		return  e
}
