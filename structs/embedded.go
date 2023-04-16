package main

import "fmt"

type AA struct {
}

func (a AA) Foo() {
	fmt.Println("aa")
}

type BB struct {
	AA
}

func (b BB) Foo() {
	fmt.Println("bb")
}

type CC struct {
	BB
}

type Interface interface {
	Foo()
}

type DD struct {
	Interface
}

func main() {
	AA{}.Foo() //aa
	BB{}.Foo() //bb
	CC{}.Foo() //bb

	DD{AA{}}.Foo() //aa
}
