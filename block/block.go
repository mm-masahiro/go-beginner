package block

import (
	"fmt"
	"math/rand"
	"time"
)

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

	rand.Seed(time.Now().UnixNano())
	// Goでは、ifのあとにこのブロックの中だけで有効な変数を定義することができる！！
	// if n := rand.Intn(10); n == 0 {
	if n := 0; n == 0 {
		fmt.Println("n  is 0")
	} else if n < 10 {
		fmt.Println("n is greater than 0 but, smaller than 10")
	} else {
		fmt.Println("n is greater than 10")	
	}
}
