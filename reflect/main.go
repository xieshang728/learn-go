package main

import (
	"fmt"
	"reflect"
)

func main(){
	x := 2
	a := reflect.ValueOf(2)
	fmt.Println(x)
	fmt.Println(a)
}