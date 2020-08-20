package main

import "fmt"

//recover 必须在defer中调用，才有效，否则返回nil
func main() {
	if info := recover(); info != nil {
		fmt.Println("catch: ", info)
	} else {
		fmt.Println("recover return nil")
	}

	defer func() {
		fmt.Println("defer one")
	}()

	defer func() {
		fmt.Println("defer two")
	}()

	defer func() {
		fmt.Println("defer three")
	}()

	panic("panic here")
}
