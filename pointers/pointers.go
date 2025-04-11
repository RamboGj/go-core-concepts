package pointers

import "fmt"

func Main() {
	x := 10
	Take(&x)
	fmt.Println(x)

	y := Create()
	fmt.Println(y)
}

func Take(x *int) {
	*x = 100
}

func Create() *int {
	x := 10
	return &x
}