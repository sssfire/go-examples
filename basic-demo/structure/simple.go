package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	bookID  int
}

func main() {

	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", bookID: 6495407})

	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})

	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.bookID = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.bookID = 6495700

	/* 打印 Book1 信息 */
	fmt.Println()
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 bookID : %d\n", Book1.bookID)

	/* 打印 Book2 信息 */
	fmt.Println()
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 bookID : %d\n", Book2.bookID)

	/* 打印 Book1 信息 */
	printBook1(Book1)                                         //传值
	fmt.Printf("Book Subject(changed) : %s\n", Book1.subject) //值传递不会修改成员变量

	/* 打印 Book2 信息 */
	printBook2(&Book2)                                        //传指针
	fmt.Printf("Book Subject(changed) : %s\n", Book2.subject) //引用传递会修改成员变量
}

/* 传递的参数为结构体 */
func printBook1(book Books) {
	fmt.Println()
	book.subject = "Changed_subject"
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book bookID : %d\n", book.bookID)
}

/* 传递的参数为结构体指针 */
func printBook2(book *Books) {
	fmt.Println()
	book.subject = "Changed_subject"
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book bookID : %d\n", book.bookID)
}
