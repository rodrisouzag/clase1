package main

import (
	"fmt"
)

type myError struct{}

func (e *myError) Error() string {
	return "error: el salario ingresado no alcanza el mínimo imponible"
}

func main() {
	err := myError{}
	salary := 200000
	if salary < 150000 {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Debe pagar impuesto")
	}

}
