package main

import "fmt"

// pings accepts channel that can only accept and word
func ping(pings chan<- string, s string) {
	pings <- s
}

func pong(pings <-chan string, pongs chan<- string) {
	pongs <- <-pings
}

func main() {

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "This is a test")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
