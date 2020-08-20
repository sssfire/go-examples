package main

import "fmt"

//延迟调用。多个defer，依次入栈，在函数即将退出时，依次出栈调用
func main() {
	defer func() {
		fmt.Println("defer one")
	}()

	defer func() {
		fmt.Println("defer two")
	}()

	//print first
	func() {
		fmt.Println("normal print")
	}()

	defer func() {
		fmt.Println("defer three")
	}()
}
