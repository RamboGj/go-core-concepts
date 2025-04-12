package main

import "fmt"

func main() {
	foo(123)
	foo("123")

	var myStruct MyStruct[string] = MyStruct[string]{}

	fmt.Println(myStruct)
}

// ~ character means my constraint is gonna accept all kind of types that are strings, as if it's not a raw string (e.g. type MyString string)
type MyConstraint interface {
	~int | ~string | ~[]int
}

func foo[T MyConstraint](arg T) {
	fmt.Println(arg)
}

type MyStruct[T any] struct {
	Foo T
}

// I CANNOT add constraints to a method (e.g. func (MyStruct[T]) foo[A any](param A) {})
func (MyStruct[T]) foo() {}

type Foo struct {}

func (Foo) do() {}

type Bar struct {}

func (Bar) do() {}

func fooBar[T interface{Foo | Bar}](arg T) {
	// Can not do it
	// arg.do()
}