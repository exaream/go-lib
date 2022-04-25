package main

import (
	"fmt"
	"time"
)

func main() {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, v := range items {
		// 参照「実用Go言語」P364

		// goroutine が起動している間にループが終わってしまうため
		// 変数 v を参照した時点で最後に実行されたループの変数しか取れない
		// それを回避するためにブロックスコープ内で代入し直す
		// v2 := v
		// go func() {
		//    fmt.Print(v2)
		// 	  // fmt.Printf("[value] v: %d, v2: %d, [address] v: %p, v2: %p\n", v, v2, &v, &v2)
		// }()

		// もしくはループの変数を引数として goroutine に渡す
		go func(v int) {
			fmt.Println(v)
		}(v)
	}
	time.Sleep(time.Second)
}
