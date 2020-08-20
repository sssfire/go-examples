package main

import (
	"fmt"
	"time"
)

/**
Channel 是可以控制读写权限的 具体如下:
go func(c chan int) { //读写均可的channel c } (a)
go func(c <- chan int) { //只读的Channel } (a)
go func(c chan <- int) {  //只写的Channel } (a)
*/

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		time.Sleep(1000 * time.Millisecond)
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c {
		fmt.Println(i)
	}
}
