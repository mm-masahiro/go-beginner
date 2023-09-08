package main

import (
	"fmt"
)

func sliceMain() {
	sliceTest()
}

func sliceTest() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	fmt.Println(y)
}
