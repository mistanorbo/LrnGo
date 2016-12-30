package main

import "fmt"

type rect struct {
	width, height int
}

// Adding a method using a pointer
func (r *rect) area() int {
	return r.width * r.height
}

// Adding a method using a type
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {

	r := rect{10, 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim", r.perim())

}
