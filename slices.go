package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println("empty array:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("setting elements:", s)
	fmt.Println("get second element:", s[1])
	fmt.Println("length:", len(s))

	s = append(s, "d")
	s = append(s, "e")
	fmt.Println("appended:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copied", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

  l = s[:3]
   fmt.Println("sl2:", l)

}
