package main

import "fmt"

const (
	catA = "A"
	catB = "B"
	catC = "C"
)

func calcularSalario(minutos float64, categoria string) float64 {

	horas := minutos / 60.0
	var salario float64

	switch categoria {
	case "A":
		salario = horas * 3000.0
		salario += salario * 0.5
	case "B":
		salario = horas * 1500.0
		salario += salario * 0.2
	case "C":
		salario = horas * 1000.0
	}

	return salario
}

func main() {
	var minutos float64 = 223.0
	var categoria string = "B"
	salario := calcularSalario(minutos, categoria)
	fmt.Printf("%.2f\n", salario)
}
