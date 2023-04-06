package main

import "fmt"

type MyErr struct {
	err error
}

func (e MyErr) Error() string {
	return "error"
}

func (e MyErr) Unwrap() error {
	return e.err
}

type UnwrapError interface {
	error
	Unwrap() error
}

func main() {
	fmt.Println("error: ", returnError() == nil)                     //true
	fmt.Println("error ptr: ", returnErrorPtr() == nil)              //true
	fmt.Println("custom error: ", returnCustomError() == nil)        //false
	fmt.Println("custom error ptr: ", returnCustomPtrError() == nil) //false
	fmt.Println("custom error nil: ", returnMyError() == nil)        //true
	fmt.Println("custom error int nil: ", returnIntError() == nil)   //false
}

func returnError() error {
	var err error
	return err
}

func returnErrorPtr() *error {
	var err *error
	return err
}

func returnCustomError() error {
	var customErr MyErr
	return customErr
}

/*
<dataPtr, baseType> мы резервируем память под переменную в dataPtr,
но по указателю лежит nil
*/
func returnCustomPtrError() error {
	var customErr *MyErr
	return customErr
}

// не возвращать указатель на ошибку!!!!
func returnMyError() *MyErr {
	return nil
}

func returnIntError() UnwrapError {
	var err *MyErr
	return err
}
