package main

import "fmt"

func printHello(ch chan int) {
	fmt.Println("hello from printHello")
	ch <- 2
}

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("Hello inline")
		ch <- 1
	}()

	go printHello(ch)

	fmt.Println("Hello from main")
	i := <-ch
	fmt.Println("Received ", i)
	<-ch
}
