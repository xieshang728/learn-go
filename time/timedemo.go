package main

import (
	"fmt"
	"time"
)

func timeDemo() {
	time := time.Now()
	year := time.Year()
	month := time.Month()
	day := time.Day()
	hour := time.Hour()
	minute := time.Minute()
	second := time.Second()
	fmt.Printf("%d-%02d-%02d %d:%d:%d\n", year, month, day, hour, minute, second)
}

func timeStamp() {
	now := time.Now()
	timestamp1 := now.Unix()
	timestamp2 := now.UTC()
	timestamp3 := now.UnixNano()

	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
	fmt.Printf("current timestamp3:%v\n", timestamp3)
}

func timeDuration() {
	duration := time.Duration(1566718614566623000)
	fmt.Println(duration.String())
	fmt.Printf("minute %s\f\n", 2*time.Minute)
}

func timer() {
	timer := time.Tick(2 * time.Second)
	for i := range timer {
		fmt.Println("timer----->", i)
	}
}
func main() {
	timeDemo()
	fmt.Println("==================")
	timeStamp()
	fmt.Println("==================")
	timeDuration()
	fmt.Println("=======timer======")
	timer()
}
