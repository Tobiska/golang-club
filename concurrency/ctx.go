package main

import (
	"fmt"
	"sort"
)

type A struct {
	i int
}

func (a A) foo() string {
	return "dsdsd"
}

type MyStruct struct {
	A
	Name string
}

func (m MyStruct) foo() string {
	return "hello"
}

func main() {
	st := []MyStruct{
		{
			Name: "Aboba",
		},
		{
			Name: "Cboba",
		},
		{
			Name: "Bebeza",
		},
	}

	fmt.Println(st[0].foo())

	sort.Slice(st, func(i, j int) bool {
		return st[i].Name < st[j].Name
	})

	fmt.Println(st)
}
