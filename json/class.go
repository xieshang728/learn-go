package main

import "fmt"

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color"`
	Actors []string
}

func main() {
	move := &Movie{}
	fmt.Println(move)
}
