/**
Go 语言切片是对数组的抽象。
Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。
*/
package main

import "fmt"

func main() {
	//定义切片
	//声明一个未指定大小的数组来定义切片，切片不需要说明长度
	var identifier []int
	printSliceLength(identifier) //len: 0, capacity: 0

	//使用make()函数来创建切片, 5是切片长度
	var slice1 []int = make([]int, 5)
	printSliceLength(slice1) //len: 5, capacity: 5

	//也可以简写为
	slice2 := make([]int, 5)
	printSliceLength(slice2) //len: 5, capacity: 5

	//也可以指定容量, 5是切片长度，3是容量
	slice3 := make([]int, 5, 10)
	printSliceLength(slice3) //len: 5, capacity: 10

	//直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3
	s := []int{1, 2, 3}
	printSliceLength(s) //len: 3, capacity: 3

	//初始化切片s,是数组arr的引用
	s = slice3[:]
	printSliceLength(s) //len: 5, capacity: 10

	startIndex := 0
	endIndex := len(slice3)

	//将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片
	s = slice3[startIndex:endIndex]
	printSliceLength(s) //len: 5, capacity: 10

	//默认 endIndex 时将表示一直到arr的最后一个元素
	s = slice3[startIndex:]

	//默认 startIndex 时将表示从arr的第一个元素开始
	s = slice3[:endIndex]
	printSliceLength(s) //len: 5, capacity: 10
}

func printSliceLength(arr []int) {
	fmt.Printf("arr length: %d, capacity: %d\n", len(arr), cap(arr))
}
