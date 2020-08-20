package main

import "fmt"

func main() {
	addFunc := add(1, 2)
	fmt.Println(addFunc())
	fmt.Println(addFunc())
	fmt.Println(addFunc())
}

// 闭包使用方法
func add(x1, x2 int) func() (int, int) {
	i := 0
	return func() (int, int) {
		i++
		return i, x1 + x2
	}
}
