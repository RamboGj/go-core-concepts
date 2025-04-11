package main

import (
	"encoding/json"
	"fmt"
)

type MinhaString string

type User struct {
	Name string `json:"name"`
	ID   uint64 `json:"id"`
}

// This adds a method for the User struct
func (u User) PrintName() {
	fmt.Println(u.Name)
}

func (u User) create() User {
	return User{Name: "Foo", ID: 10101}
}

func (u *User) UpdateName(newName string) {
	u.Name = newName
}

func main() {
	user := User{
		Name: "User name",
		ID:   3123132,
	}
	user.UpdateName("Joaquim")
	fmt.Println(user)

	jsonUser, err := json.Marshal(user)

	if err != nil {
		panic("err")
	}

	fmt.Println(string(jsonUser))

	// newUser := user.create()
	// user.PrintName()
	// fmt.Println(newUser)
}