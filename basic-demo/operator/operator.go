package main

import "fmt"

func main() {

	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c)
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c)
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c)
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c)
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a)
	a = 21 // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a)

	fmt.Println("=============================")
	a = 21
	b = 10

	if a == b {
		fmt.Printf("第一行 - a 等于 b\n")
	} else {
		fmt.Printf("第一行 - a 不等于 b\n")
	}
	if a < b {
		fmt.Printf("第二行 - a 小于 b\n")
	} else {
		fmt.Printf("第二行 - a 不小于 b\n")
	}

	if a > b {
		fmt.Printf("第三行 - a 大于 b\n")
	} else {
		fmt.Printf("第三行 - a 不大于 b\n")
	}
	/* Lets change value of a and b */
	a = 5
	b = 20
	if a <= b {
		fmt.Printf("第四行 - a 小于等于 b\n")
	}
	if b >= a {
		fmt.Printf("第五行 - b 大于等于 a\n")
	}

	fmt.Println("=============================================")
	var e bool = true
	var f bool = false
	if e && f {
		fmt.Printf("第一行 - 条件为 true\n")
	}
	if e || f {
		fmt.Printf("第二行 - 条件为 true\n")
	}

	e = false
	f = true
	if e && f {
		fmt.Printf("第三行 - 条件为 true\n")
	} else {
		fmt.Printf("第三行 - 条件为 false\n")
	}
	if !(e && f) {
		fmt.Printf("第四行 - 条件为 true\n")
	}

	fmt.Println("==================================")
	var g int = 4
	var h int32
	var i float32
	var ptr *int

	/* 运算符实例 */
	fmt.Printf("第 1 行 - g 变量类型为 = %T\n", g)
	fmt.Printf("第 2 行 - h 变量类型为 = %T\n", h)
	fmt.Printf("第 3 行 - i 变量类型为 = %T\n", i)

	/*  & 和 * 运算符实例 */
	ptr = &g /* 'ptr' 包含了 'g' 变量的地址 */
	fmt.Printf("g 的值为  %d\n", g)
	fmt.Printf("*ptr 为 %d\n", *ptr)
}
