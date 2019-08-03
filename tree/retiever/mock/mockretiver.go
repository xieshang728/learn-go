package mock

import "fmt"

type Retriever struct {
	Contents string
}

func (r Retriever) Get(url string) string{
	fmt.Println("mock get params "+url)
	return r.Contents
}
