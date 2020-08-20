package main

import "fmt"

func main() {

	var a int
	var ptr *int
	var pptr **int

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)

	var a1 int = 5
	//把ptr指针 指向ss所在地址
	var ptr1 *int = &a1
	//开辟一个新的指针，指向ptr指针指向的地方
	var pts1 *int = ptr1
	//二级指针，指向一个地址，这个地址存储的是一级指针的地址
	var pto1 **int = &ptr1
	//三级指针，指向一个地址，这个地址存储的是二级指针的地址，二级指针同上
	var pt31 ***int = &pto1
	fmt.Println("a的地址:", &a1,
		"\n 值", a1, "\n\n",

		"ptr指针所在地址:", &ptr1,
		"\n ptr指向的地址:", ptr1,
		"\n ptr指针指向地址对应的值", *ptr1, "\n\n",

		"pts指针所在地址:", &pts1,
		"\n pts指向的地址:", pts1,
		"\n pts指针指向地址对应的值:", *pts1, "\n\n",

		"pto指针所在地址:", &pto1,
		"\n pto指向的指针（ptr）的存储地址:", pto1,
		"\n pto指向的指针（ptr）所指向的地址: ", *pto1,
		"\n pto最终指向的地址对应的值（a）", **pto1, "\n\n",

		"pt3指针所在的地址:", &pt31,
		"\n pt3指向的指针（pto）的地址:", pt31, //等于&*pt31,
		"\n pt3指向的指针（pto）所指向的指针的（ptr）地址", *pt31, //等于&**pt31,
		"\n pt3指向的指针（pto）所指向的指针（ptr）所指向的地址（a）:", **pt31, //等于&***pt31,
		"\n pt3最终指向的地址对应的值（a）", ***pt31)

}
