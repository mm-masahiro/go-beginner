package main

import "fmt"

const x int64 = 10

const (
	idKey = "id"
	nameKey = "name"
)

const z = 20 * 10

func main() {
  const y = "Hello"

	fmt.Println(x)
	fmt.Println(y)

	fmt.Println(x)
	fmt.Println(y)
	
	var a [2][3]int

	fmt.Println(a)
}

