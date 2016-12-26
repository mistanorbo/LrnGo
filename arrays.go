package main

import "fmt"

func main() {
  var a [5]int
  fmt.Println(a)
  a[4] = 10
  fmt.Println(a)
  fmt.Println(a[4])

  b := [5]int{1,2,3,4,5}
  fmt.Println(b)
}
