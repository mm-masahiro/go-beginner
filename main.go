package main

import (
	"fmt"
)

type person struct {
	age  int
	name string
}

func modifyFails(i int, s string, p person) {
	i *= 2
	s = "good bye"
	p.name = "Bob"
	fmt.Println(p)
}

func modMap(m map[int]string) {
	m[2] = "Hello"
	m[3] = "Good Bye"

	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}

	s = append(s, 10)
}

func main() {
	a := "aaaa"
	fmt.Println(&a)

	var address_a *string = &a

	// ポインタ変数から値を間接参照
	fmt.Println(*address_a)

	// ポインタ変数を通して、変数aの値を書き換え
	*address_a = "bbbb"

	fmt.Println(*address_a)
	fmt.Println(a)

	// defer fmt.Println("exec1 in main")
	// defers.Defer()
	// defers.MultiDefer()

	// p := person{}
	// i := 2
	// s := "Hello"
	// fmt.Println(i, s, p)
	// modifyFails(i, s, p)
	// fmt.Println(i, s, p)

	m := map[int]string{
		1: "first",
		2: "second",
	}

	modMap(m)
	fmt.Println(m)

	s := []int{1, 2, 3}
	modSlice(s)
	fmt.Println(s)
}
