//panic和defer结合使用：panic触发错误，defer依次出栈调用，没有recover捕获的情况下，最后才打印错误
package main

import "fmt"

//panic 其后的代码不会执行
func main() {
	defer func() {
		fmt.Println("defer one")
	}()

	defer func() {
		fmt.Println("defer two")
	}()

	//print first
	fmt.Println("normal print")

	defer func() {
		fmt.Println("defer three")
	}()

	panic("panic here")

	//cannot reach here and will not print
	fmt.Println("normal print last")
}
