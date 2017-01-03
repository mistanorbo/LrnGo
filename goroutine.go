package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")

	go func(msg string) {
		time.Sleep(2 * time.Second)
		fmt.Println(msg)
	}("going")

	go f("goroutine")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")

}
