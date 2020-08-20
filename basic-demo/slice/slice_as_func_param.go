package main

import "fmt"

func main() {
	changeSliceTest()
}

func changeSliceTest() {
	arr1 := []int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{1, 2, 3}

	fmt.Println("before change arr1, ", arr1)
	changeSlice(arr1) // slice 按引用传递
	fmt.Println("after change arr1, ", arr1)

	fmt.Println("before change arr2, ", arr2)
	changeArray(arr2) // array 按值传递
	fmt.Println("after change arr2, ", arr2)

	fmt.Println("before change arr3, ", arr3)
	changeArrayByPointer(&arr3) // 可以显式取array的 指针
	fmt.Println("after change arr3, ", arr3)
}

func changeSlice(arr []int) {
	arr[0] = 9999
}

func changeArray(arr [3]int) {
	arr[0] = 6666
}

func changeArrayByPointer(arr *[3]int) {
	arr[0] = 6666
}
