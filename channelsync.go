package main

import "fmt"
import "time"

func worker(done chan bool) {
	fmt.Println("Doing something...")
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
	done <- true
}

func main() {

	done := make(chan bool, 1)
	go worker(done)

	<-done
}
