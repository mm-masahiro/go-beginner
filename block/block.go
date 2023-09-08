package block

import "fmt"

func Block() {
	x := 10

	if x > 5 {
		fmt.Println(x)
		// -> printed 10
		x := 5
		fmt.Println(x)
		// -> printed 5
	}

	fmt.Println(x)
	// -> printed 10
}
