package main

import (
	"fmt"
	function "go-beginner/function"
)

type person struct {
	age  int
	name string
}

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

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
	// const y = "Hello"

	// fmt.Println(x)
	// fmt.Println(y)

	// fmt.Println(x)
	// fmt.Println(y)

	// var a [2][3]int

	// fmt.Println(a)

	// var s = []int{10,20,30}
	// fmt.Println(s)
	// s = append(s, 40)
	// fmt.Println(s)
	// s = append(s, s...)
	// fmt.Println(s)

	// sWithCap := make([]int, 5)
	// fmt.Println(sWithCap)

	// var nilMap map[string]int
	// fmt.Println(nilMap["a"])

	// teams := map[string][]string {
	// 	"team1": []string{"hoge", "foo" },
	// 	"team2": []string{"ho", "foooooo"},
	// }
	// fmt.Println(teams["team1"])

	// delete(teams, "team2")
	// fmt.Println(teams)

	// type person struct {
	// 	name string
	// 	age int
	// 	pet string
	// }

	// fred := person{
	// 	name: "fred",
	// 	age: 23,
	// 	pet: "dog",
	// }

	// pet := struct {
	// 	name string
	// 	age int
	// } {
	// 	name: "hachi",
	// 	age: 11,
	// }

	// fmt.Println(pet)

	// fmt.Println(fred)
	// bl.Block()
	// block.For()
	// block.While()
	// block.ForRange()

	function.MyFunc(function.FuncOptions{
		LastName: "Hoge",
		Age:      20,
	})

	function.AddTo(3)
	function.AddTo(3, 8)
	function.AddTo(3, 8, 10)
	function.AddTo(3, 8, 10, 23)

	a := "aaaa"
	fmt.Println(&a)

	var address_a *string = &a

	// ポインタ変数から値を間接参照
	fmt.Println(*address_a)

	// ポインタ変数を通して、変数aの値を書き換え
	*address_a = "bbbb"

	fmt.Println(*address_a)
	fmt.Println(a)

	// str.Structures()

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
