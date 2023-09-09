package block

import "fmt"

func For() {
  for i := 0; i < 10; i++ {
    if i == 0 {
      fmt.Println("i is 0,", i)
      i := 100
      fmt.Println("i is overrided!!", i)
    }
    fmt.Println(i)  
  } 
}

func While() {
  i := 1
  for i < 100 {
    fmt.Println(i)

    i = i * 2
  }
}
