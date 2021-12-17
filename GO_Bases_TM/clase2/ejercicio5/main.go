package main

import (
	"errors"
	"fmt"
)

const (
	perro = "perro"
	gato  = "gato"
)

func animalPerro(cantidad float64) float64 {
	return cantidad * 10.0
}

func animalGato(cantidad float64) float64 {
	return cantidad * 5.0
}

func animalHamster(cantidad float64) float64 {
	return cantidad * 0.25
}

func animalTarantula(cantidad float64) float64 {
	return cantidad * 0.15
}

func Animal(tipo string) (func(cantidad float64) float64, error) {
	switch tipo {
	case "perro":
		return animalPerro, nil
	case "gato":
		return animalGato, nil
	case "hamster":
		return animalHamster, nil
	case "tarantula":
		return animalTarantula, nil
	default:
		return nil, errors.New("No existe este tipo de animal")
	}
}

func main() {
	var cantidad float64 = 0.0

	animalPerro, err := Animal(perro)
	animalGato, err := Animal(gato)

	if err == nil {
		cantidad += animalPerro(5)
		cantidad += animalGato(8)
		cantidad += animalTarantula(17)
		fmt.Printf("%.2f\n", cantidad)
	} else {
		fmt.Println(err)
	}

}
