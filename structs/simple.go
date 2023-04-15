package main

import "fmt"

type A struct{}

type B struct {
	A    //композиция
	a *A //агрегация
}

type MyStruct struct {
	wh, pt int
	name   string
}

func (m *MyStruct) String() string {
	return fmt.Sprintf(`
		wh: %d,
		pt: %d,
		name: %s 
	`, m.wh, m.pt, m.name)
}

func main() {
	fmt.Println(MyStruct{name: "Pipa"})
}
