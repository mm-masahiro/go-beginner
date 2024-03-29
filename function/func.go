package function

import "fmt"

func Div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}

	return numerator / denominator
}

type FuncOptions struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(options FuncOptions) error {
	fmt.Println(options)

	return nil
}

func AddTo(base int, values ...int) []int {
	out := make([]int, 0, len(values))
	for _, v := range values {
		out = append(out, base+v)
	}

	return out
}
