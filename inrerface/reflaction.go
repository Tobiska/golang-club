package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 1. Reflection распространяется от интерфейса до reflection объекта.
	//
	//На базовом уровне reflect является всего лишь механизмом для изучения пары тип и значение,
	//хранящейся внутри переменной интерфейса
	//
	//reflect.Type и reflect.Value

	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x)) // type: float64  принимает - interface{}

	//Когда мы вызываем reflect.TypeOf(x), x сначала сохраняется в пустом интерфейсе,
	//который затем передается в качестве аргумента;
	//reflect.TypeOf() распаковывает этот пустой интерфейс для восстановления информации о типе

	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	//Метод Kind(), возвращает константу, указывающую, какой примитивный элемент хранится: Uint, Float64, Slice…
	//Эти константы объявлены в перечислении в пакете reflect
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	// 2. Reflection распространяется от reflect объекта до интерфейса.

	// Имея reflect.Value, мы можем восстановить значение интерфейса с помощью метода Interface();
	// метод упаковывает информацию о типе и значении обратно в интерфейс и возвращает результат:
	// func (v Value) Interface() interface{}

	var xdd float64 = 3.433 //
	ddd := reflect.ValueOf(xdd)
	y := ddd.Interface().(float64) // y имеет тип float64.
	fmt.Println(y)

	// 3. Чтобы изменить объект отражения, значение должно быть устанавливаемым.

	// Устанавливаемость немного напоминает адресуемость, но строже.
	// Это свойство, при котором reflection объект может изменить хранимое значение,
	// которое было использовано при создании reflection объекта.
	// Устанавливаемость определяется тем, содержит ли reflection объект исходный элемент, или только его копию!

	var x3 float64 = 3.4
	v4 := reflect.ValueOf(x3)
	v4.SetFloat(7.1) //  panic: reflect.Value.SetFloat использует неадресуемое значение

	fmt.Println("settability of v:", v.CanSet()) // false

	var x4 float64 = 3.4
	p := reflect.ValueOf(&x4) // Берём адрес x. и поэтому можем изменить основной обьек с помощью рефлексии.
	// Это сделано для избежания путаницы
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet()) // false  т.к. это указатель *p поэтому мы и не можем его установить

	v6 := p.Elem()                                // получение p
	fmt.Println("settability of v:", v6.CanSet()) // true
}
