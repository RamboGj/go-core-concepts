package main

import "fmt"

// Interfaces -> defines the methods a type needs to have to impl it

type Animal interface {
	Sound() string
}

type Dog struct {}

func (*Dog) Sound() string {
	return "Au! Au!"
}

type Cat struct {}

func (*Cat) Sound() string {
	return "Miau!"
}

func whatDoesThisAnimalSay(a Animal) {
	fmt.Println(a.Sound())
} 

func main() {
	dog := Dog{}
	whatDoesThisAnimalSay(&dog)
}

func foo (a interface{}) {}

// Type assertion -> verify a variable type
func takeAssertion(a any) {
	str, ok := a.(string)

	if ok {
		fmt.Println("a is an string", str)
	}
}

// Switch types
func takeSwitchTypes(a Animal) {
	switch t := a.(type) {
	case *Dog:
		t.Sound()
	case *Cat:	
		t.Sound()
	}
	
}