package main

import (
	"bufio"
	"github.com/gpmgo/gopm/modules/log"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"io"
	"learn-go/crawler/engine"
	"learn-go/crawler/zhenai/parser"
)

func main(){

	engine.Run(engine.Request{Url:"http://www.zhenai.com/zhenghun",ParseFunc:parser.ParseCityList})

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