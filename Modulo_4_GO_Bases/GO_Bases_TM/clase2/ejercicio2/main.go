package main

import (
	"errors"
	"fmt"
)

func calcularPromedio(calificaciones ...int) (int, error) {
	var suma int = 0
	for _, cal := range calificaciones {
		if cal < 0 {
			return 0, errors.New("Ha ingresado una calificacion negativa")
		} else {
			suma += cal
		}
	}
	return suma / len(calificaciones), nil
}

func main() {
	promedio, err := calcularPromedio(12, 10, 5, 7, 9)
	if err == nil {
		fmt.Println(promedio)
	} else {
		fmt.Println(err)
	}
}
