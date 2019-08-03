package main

import "fmt"

func main(){
	natures := make(chan int)
	squares := make(chan int)

	go func(){
		for x := 0; ; x++{
			natures <- x
		}
	}()

	go func() {
		for {
			x := <- natures
			squares <- x*x
		}
	}()

	fmt.Println(<- squares)
}
