package main

import "fmt"

// Пустой interface{}
// Этому интерфейсу удовлетворяет вообще любой тип.
// Поскольку у пустого интерфейса нет никаких методов, то и itable для него просчитывать и хранить не нужно
// — достаточно только метаинформации о статическом типе.
// Пустой интерфейс это интерфейс ккоторому удовлетворяет любой тип

// empty interface as function parameter
func displayValue(i interface{}) {
	fmt.Println(i)
}

func main() {
	// variable of empty interface type
	var a interface{}
	fmt.Println("Value:", a) //Value: <nil>

	b := 20
	c := "Test"
	// pass integer number to the function
	displayValue(b)
	// pass boolean value to the function
	displayValue(c)

}
