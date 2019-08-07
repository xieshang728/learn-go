package engine

import (
	"learn-go/crawler/fetcher"
	"log"
)

func Run(sends ...Request){
	var requests []Request

	for _,r := range sends{
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests := requests[1:]
		body,err := fetcher.Fetch(r.Url)
		if err != nil{
			log.Printf("Fetcher: error fetching url %s: %v",r.Url,err)
			continue
		}
		parseResult := r.ParseFunc(body)
		requests = append(requests,parseResult.Requests...)

		for _,item := range parseResult.Items{
			log.Printf("Got item %v",item)
		}
	}
}
