package main

import (
	"fmt"
	"time"
)

/**
	从结果可以看出，world的打印不到5个就结束了，原因是say1还未执行完毕主线程就结束了，从而say1被结束之行
	如果想要阻止主函数结束，要等待goroutine结束后，再退出主函数，可以查看注释掉的代码
**/
func say1(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func say2(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say1("world")
	say2("hello")
}

// func say1(s string, ch chan int) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// 	ch <- 0
// 	close(ch)
// }

// func main() {
// 	ch := make(chan int)
// 	go say1("world", ch)
// 	say2("hello")
// 	fmt.Println(<-ch)
// }
