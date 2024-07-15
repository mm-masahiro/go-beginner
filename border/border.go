package border

import "fmt"

type bord interface {
	lower_layer() int
}

type str struct{}

func Border() {
	fmt.Println(bad_upper_layer())
	// fmt.Println(good_upper_layer(str{}))
}

func bad_upper_layer() string {
	r := lower_layer()
	return fmt.Sprintf("lower is %d", r)
}

// interfaceを引数に入れることで、下位層への依存をなくした(?)
// 何が嬉しいのか考える
func good_upper_layer(b bord) string {
	r := b.lower_layer()
	return fmt.Sprintf("lower is %d", r)
}

func lower_layer() string {
	a := 1
	b := 2
	return fmt.Sprint("a is %d, b is %d", a, b)
}

func (s str) lower_layer() string {
	a := 1
	b := 2
	return fmt.Sprint("a is %d, b is %d", a, b)
}
