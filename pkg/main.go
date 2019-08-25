package main

import (
	crand "crypto/rand"
	"fmt"
	"math/rand"
)

func main(){
	fmt.Println(rand.Int())
	fmt.Println(crand.Read([]byte{}))
}
