package main

import "fmt"

//MAX define a constant
const MAX int = 3

func main() {
	a := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}

	//创建指针数组的时候，不适合用 range 循环
	number := [MAX]int{5, 6, 7}
	var ptrs [MAX]*int //指针数组
	//将number数组的值的地址赋给ptrs
	for i := range number {
		ptrs[i] = &number[i]
	}
	for i, x := range ptrs {
		fmt.Printf("指针数组：索引:%d 值:%d 值的内存地址:%d\n", i, *x, x)
	}
}
