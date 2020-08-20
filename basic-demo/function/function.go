package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	var ret int

	/* 调用函数并返回最大值 */
	ret = max(a, b)

	fmt.Printf("最大值是 : %d\n", ret)

	/* 返回多个值 */
	a1, b1 := swap("Google", "Runoob")
	fmt.Println(a1, b1)

	/* 值引用 */
	a2, b2 := "google", "runoob"
	swap(a2, b2)
	fmt.Println("a, b: ", a2, b2)
	fmt.Println(a2, b2)

}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

/* 返回多个值 */
func swap(x, y string) (string, string) {
	return y, x
}

/* 引用传递 */
func swap1(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}
