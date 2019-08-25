package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("1 号协程完成")
		wg.Done()
	}()

	go func() {
		time.Sleep(4 * time.Second)
		fmt.Println("2 号协程完成")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("所有协程完成")
}
