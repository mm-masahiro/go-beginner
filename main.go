package main

import (
	function "go-beginner/function"
)

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

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
}
