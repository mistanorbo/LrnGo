package main

import "fmt"

func zeroval(ival int) {
	ival = 0
	fmt.Println("zeroval-internal:", ival)
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {

	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	// https://gobyexample.com/pointers

	// http://piotrzurek.net/2013/09/20/pointers-in-go.html

}
