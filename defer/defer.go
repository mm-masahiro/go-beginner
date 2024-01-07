package defers

import (
	"context"
	"database/sql"
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

func DoSomeInserts(ctx context.Context, db *sql.DB, value1, value2 string) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		// deferはreturnされた後に実行されるので、名前付き戻り値との比較がクロージャの中でできる
		if err == nil {
			err = tx.Commit()
		}
		if err != nil {
			tx.Rollback()
		}
	}()
	_, err = tx.ExecContext(ctx, "INSERT INTO FOO (val) values $1", value1)
	if err != nil {
		return err
	}

	return nil
}
