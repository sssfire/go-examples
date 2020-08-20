package main

import "fmt"

//defer,panic, recover 结合使用，panic触发错误，defer依次出栈调用，直到被recover捕获，打印捕获的信息，之后继续defer出栈
func main() {
	defer func() {
		fmt.Println("defer one")
	}()

	defer func() {
		if info := recover(); info != nil {
			fmt.Println("catch: ", info)
		}
		fmt.Println("defer three")
	}()

	defer func() {
		fmt.Println("defer two")
	}()

	panic("panic here")
}
