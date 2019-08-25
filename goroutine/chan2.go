package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控停止了..............")
				return
			default:
				fmt.Println("监控中.................")
				time.Sleep(2 * time.Second)
			}
		}
	}()
	time.Sleep(10 * time.Second)
	fmt.Println("可以通知停止了...........")
	stop <- true
	time.Sleep(5 * time.Second)
}
