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

	var s = []int{10,20,30}
	fmt.Println(s)
	s = append(s, 40)
	fmt.Println(s)
	s = append(s, s...)
	fmt.Println(s)

	sWithCap := make([]int, 5)
	fmt.Println(sWithCap)
}

