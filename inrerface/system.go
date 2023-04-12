package main

import "unsafe"

//https://research.swtch.com/interfaces

//src/runtime/runtime2.go

type iface struct {
	tab *itab // указатель на Interface Table или itable — структуру,
	// которая хранит некоторые метаданные о типе и список методов,
	//используемых для удовлетворения интерфейса.
	data unsafe.Pointer // указывает на фактическую переменную с конкретным (статическим) типом
}

type itab struct {
	inter *interfacetype
	//_type *_type
	hash uint32 // copy of _type.hash. Used for type switches.
	_    [4]byte
	fun  [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
}

type interfacetype struct {
	//typ     _type
	//pkgpath name
	//mhdr    []imethod
}
