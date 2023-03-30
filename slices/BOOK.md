## Массивы

Массив - типизированная коллекция фиксированного размера.

Комбинация размера и типа элемента образуют новый тип.

```golang
arr := [5]int{5, 7, 1, 2, 0}
``` 

Массив хранится в виде последовательности из n блоков определенного типа.

```golang
func newarray(typ *_type, n int) unsafe.Pointer {
	if n == 1 {
		return mallocgc(typ.size, typ, true)
	}
	mem, overflow := math.MulUintptr(typ.size, uintptr(n)) //вычисление кол-во памяти, смотрим если умножение переполняется
	if overflow || mem > maxAlloc || n < 0 {
		panic(plainError("runtime: allocation size out of range"))
	}
	return mallocgc(mem, typ, true)
}
```
Ref: https://go.dev/src/runtime/malloc.go

## Слайсы

**Слайс** - модифицированный указатель на базовый массив(необязательно на первый элемент!).

```golang
type slice struct {
	array unsafe.Pointer // указатель на первый элемент массива
	len   int // фактическое кол-во элементов
	cap   int // ёмкость слайса
}
```

![alt text](https://drstearns.github.io/tutorials/goslicemap/img/slices-2.png)

**Слайс**, в отличии от массива, не выделяет память во время инициализации.
Фактически, срезы инициализируется с нулевым (nil) значением.

### Runtime

Ref: https://go.dev/src/runtime/slice.go

#### **makeslicecopy**
#### **makeslice**
**makeslice** используется при **make(slice, len, cap)**, если **cap** отсутствует, по умолчанию он будет таким же, как len.
 - Попробуйте аллоцировать память размера **cap**.
 - Если **len** > **cap**, выделить память длины **len**.
 - Если **len < 0** или слишком много, паникуем.
#### **growslice**
 - по умл. ставится 2*oldcap
 - если newcap > 2*oldcap, то newcap
 - если newcap > 256, то увеличиваем в 1.25 раза
#### **slicecopy**
slicecopy используется в двух случаях:
 - После **growslice** для копирования старого массива.
 - Функия **copy()**.
 
Обратите внимание, что slicecopy гарантирует только memmove, а не пытается выделить память.




