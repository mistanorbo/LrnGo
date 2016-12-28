package main

import "fmt"

func add(a int, b int) int {

  return a + b

}

func addAll(nums ...int) int {

  total := 0

  for _, num := range nums {

    total += num

  }

  return total
}

func main() {

  fmt.Println("result:", addAll(1,2,3,4,5))

}
