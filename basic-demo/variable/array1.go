package main

import "fmt"

func main() {
	var n [10]int /* n 是一个长度为 10 的数组 */
	var i, j int
	var balance1 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}   //设置数组大小
	var balance2 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0} //不设置数组大小

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	for i = 0; i < len(balance1); i++ {
		fmt.Printf("balance1[%d] = %f\n", i, balance1[i])
	}

	for i = 0; i < len(balance2); i++ {
		fmt.Printf("balance2[%d] = %f\n", i, balance2[i])
	}
}
