package main

import (
	"fmt"
	"time"
)

func createChannel(i int) chan <- int{
	c := make(chan int)

	go func() {
		for {
			n := <- c
			fmt.Printf("work %d ,received %c\n",i,n)
		}
	}()

	return c
}

func channelDemo(){

	var channels [10] chan <- int

	for i := 0; i < 10; i++{
		channels[i] = createChannel(i)
	}

	for i:=0; i < 10; i++  {
		channels[i] <- i+'a'
	}

	for i:=0; i < 10; i++ {
		channels[i] <- i +'A'
	}

	time.Sleep(time.Millisecond)
}

func bufferChan(){
	c := make(chan int,3)
	c <- 1
	c <- 2
	c <- 3
}

func main(){
	//channelDemo()
	bufferChan()
}
