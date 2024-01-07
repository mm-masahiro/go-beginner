package defers

import (
	"fmt"
	"io"
	"log"
	"os"
)

func MultiDefer() {
	fmt.Println("exec1 in multi")
	// deferはLIFOなので、後に定義されたもののほうが先に実行される
	defer fmt.Println("exec4 in multi")
	defer fmt.Println("exec3 in multi")
	defer fmt.Println("exec2 in multi")

}

func Defer() {
	if len(os.Args) < 2 {
		log.Fatal("not selected file")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close() // clean upするコード

	data := make([]byte, 2048)

	for {
		// count・・・読み込んだbyte数
		count, err := f.Read(data)
		// 標準出力に出力
		os.Stdout.Write(data[:count])

		if err != nil {
			// ファイルの終わりでないときのエラー処理
			if err != io.EOF {
				log.Fatal(err)
			}
		}
	}
}
