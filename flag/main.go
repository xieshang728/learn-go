package main

import (
	"flag"
	"fmt"
	"os"
)

var period = flag.String("period", "default", "default")

func main() {
	allArgs := os.Args
	args := os.Args[1:]

	fmt.Println(allArgs)
	fmt.Println(args)

	fmt.Println("===================================")
	flag.Parse()
	fmt.Println(*period)
}
