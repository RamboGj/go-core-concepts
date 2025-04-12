package errors

import (
	"errors"
	"fmt"
	"math"
)

/*
Errors -> standard language interfaces
*/

type SqrtError struct {
	msg string
}

func (se SqrtError) Error() string { return se.msg }

var ErrNotFound = errors.New("Not found")

func sqrtRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, SqrtError{"There is no square root for a negative number"}
	}

	result := math.Sqrt(x)
	return result, nil
}

func main() {
	// res, err := sqrtRoot(-10)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(res)

	err := foo()

	if err != nil && errors.Is(err, ErrNotFound) {
		fmt.Println("Is not found error")
		return 
	} 

	// OR
	err2 := foo()
	var sqrtError SqrtError

	if err2 != nil && errors.As(err2, &sqrtError /*We pass it as a pointer because error.As() function needs to mutate the variable in case it matchs, so the variable err2 has the SqrtError properties*/) {
		fmt.Println("Is SqrtError")
	}

}

func foo() error { return nil } 

/* Error Wrapping */