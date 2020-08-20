package main

import (
	"fmt"
)

type fruit interface {
	getName() string
	setName(name string)
}

type apple struct {
	name string
}

//[1]
func (a *apple) getName() string {
	return a.name
}

//[2]
func (a *apple) setName(name string) {
	a.name = name
}

func main() {
	a := apple{"红富士"}
	fmt.Println(a.getName())
	a.setName("树顶红")
	fmt.Println(a.getName())	
}
