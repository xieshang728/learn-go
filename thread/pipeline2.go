package main

import "fmt"

func main(){
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 10; x++{
			fmt.Printf("naturals %d\n",x)
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals{
			fmt.Printf("squares %d\n",x)
			squares <- x
		}
		close(squares)
	}()

	for x := range squares{
		fmt.Printf("main %d\n",x)
		fmt.Println(x)
	}
}
