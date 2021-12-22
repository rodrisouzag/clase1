package main

import (
	"errors"
	"fmt"
)

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func minFunc(valores ...int) int {
	min := valores[0]
	for _, val := range valores {
		if val < min {
			min = val
		}
	}
	return min
}

func maxFunc(valores ...int) int {
	max := valores[0]
	for _, val := range valores {
		if val > max {
			max = val
		}
	}
	return max
}

func promFunc(valores ...int) int {
	sum := 0
	for _, val := range valores {
		sum += val
	}
	return sum / len(valores)
}

func operacion(tipo string) (func(valores ...int) int, error) {
	switch tipo {
	case "minimo":
		return minFunc, nil
	case "maximo":
		return maxFunc, nil
	case "promedio":
		return promFunc, nil
	default:
		return nil, errors.New("No existe el tipo de funcion ingresado")
	}
}

func main() {
	minFunc, err := operacion(minimo)
	promFunc, err := operacion(promedio)
	maxFunc, err := operacion(maximo)

	if err == nil {
		valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Printf("Minimo = %d\n", valorMinimo)

		valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Printf("Promedio = %d\n", valorPromedio)

		valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Printf("Maximo = %d\n", valorMaximo)
	} else {
		fmt.Println(err)
	}

}
