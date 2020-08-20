package main

import "fmt"

func main() {
	// declare string
	var a string = "Runoob"
	fmt.Println(a)

	// declare int var with initial value
	var b, c int = 1, 2
	fmt.Println(b, c)

	// declare a new string var with init value
	f := "Runoob" // var f string = "Runoob"
	fmt.Println(f)

	// declare multiple vars
	var vname1, vname2, vname3 int
	vname1, vname2, vname3 = 1, 2, 3
	fmt.Println(vname1, vname2, vname3)

	// declare multiple vars
	var vname4, vname5, vname6 = 1, 2, 3 // 和 python 很像,不需要显示声明类型，自动推断
	fmt.Println(vname4, vname5, vname6)

	// declare multiple vars
	vname7, vname8, vname9 := 1, 2, 3 // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误
	fmt.Println(vname7, vname8, vname9)

	// 这种因式分解关键字的写法一般用于声明全局变量
	var (
		vname10 int = 0
		vname11 int = 1
	)
	fmt.Println(vname10, vname11)

	// define constant var
	const constABC string = "abc"
	const constDEF = "def"
	fmt.Println(constABC)
	fmt.Println(constDEF)

	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)
	fmt.Println(Unknown, Female, Male)

	const (
		za = 3    // 3
		zb        // 3
		zc        // 3
		zd = "ha" // "ha"
		ze        // "ha"
		zf        // "ha"
	)
	fmt.Println(za, zb, zc, zd, ze, zf)

	const (
		aa = iota // 0
		ab        // 1
		ac        // 2
		ad = "ha" // 独立值, iota += 1
		ae        // "ha",  iota += 1
		af = 100  // iota += 1
		ag        // 100, iota += 1
		ah = iota // 7, 恢复计数
		ai        // 8
	)
	fmt.Println(aa, ab, ac, ad, ae, af, ag, ah, ai)

	const (
		bi = 1 << iota // 左移0位，不变仍为1
		bj = 3 << iota // 左移1位，11 -> 110, 6
		bk             // 左移2位，11 -> 1100, 12
		bl             // 左移3位，11 -> 11000, 24
	)
	fmt.Println("i=", bi)
	fmt.Println("j=", bj)
	fmt.Println("k=", bk)
	fmt.Println("l=", bl)
}
