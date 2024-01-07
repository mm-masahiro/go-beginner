package structures

import "fmt"

type Person struct {
	name   string
	age    int
	gender string
}

var p = Person{
	name:   "person1",
	age:    23,
	gender: "men",
}

func Structures() {
	fmt.Println(p)
}
