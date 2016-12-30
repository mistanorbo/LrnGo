package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextVal := intSeq()

	fmt.Println(nextVal())
	fmt.Println(nextVal())
	fmt.Println(nextVal())

	newVal := intSeq()

	fmt.Println(newVal())

}
