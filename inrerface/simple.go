package main

import (
	"fmt"
	"io"
	"os"
)

// Speaker  - в Go же принято использовать постфикс -er (Sender, Reader, Closer, etc)
type Speaker interface {
	SayHello()
	// age string (ошибка) — интерфейс не может содержать данные, только методы.
	//Структура хранит данные, но не поведение. Интерфейс хранит поведение, но не данные.
}

// Human - имплементация интерфейса указывается неявно(duck typing)
// В Go структура с методами будет удовлетворять интерфейсу просто самим фактом объявления метода.
type Human struct {
	Hi string
}

// SayHello - в Go — case первой буквы имени методы определяет видимость
func (h Human) SayHello() {
	fmt.Println(h.Hi)
}

func main() {
	//var s Speaker
	//s = Human{Hi: "Hello"}
	//s.SayHello()

	h := Human{Hi: "Hello"}
	s := Speaker(h)
	s.SayHello()
	//
	h.Hi = "Meow"
	s.SayHello() // выведет "Hello" т.к в переменную data копия старого Human

	//Теперь несколько слов про itable. Поскольку эта таблица будет уникальна для каждой пары интерфейс-статический тип,
	//то просчитывать её на этапе компиляции (early binding) будет нерационально и неэффективно
	//Вместо этого, компилятор генерирует метаданные для каждого статического типа, в которых, помимо прочего,
	//хранится список методов, реализованных для данного типа. Аналогично генерируются метаданные со списком методов
	//для каждого интерфейса. Теперь, во время исполнения программы, runtime Go может вычислить itable на лету (late binding)
	//для каждой конкретной пары. Этот itable кешируется, поэтому просчёт происходит только один раз.

	//var s Speaker = string("test") // compile-time error
	//var s Speaker = io.Reader // compile time error
	//var h string = Human{} // compile time error
	//var s interface{}; h := s.(Human) // runtime error
	//h.SayHello()

	var r io.Reader
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println(err)
	}

	r = tty

	var w io.Writer

	//Выражение в этом присваивании является утверждением типа;
	//оно утверждает, что элемент внутри r также реализует io.Writer,
	//и поэтому мы можем назначить его w
	w = r.(io.Writer)

	var empty interface{}
	empty = w

	fmt.Println(empty)
}
