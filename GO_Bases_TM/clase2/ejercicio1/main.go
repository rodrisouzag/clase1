package main

import (
	"fmt"
)

func calcularImpuesto(salario float64) (impuesto float64) {
	if salario > 50000.0 && salario < 100000.0 {
		impuesto = salario * (17.0 / 100.0)
	} else if salario > 100000.0 {
		impuesto = salario * (27.0 / 100.0)
	} else {
		impuesto = 0.0
	}
	return
}

func main() {

	var salario float64 = 53456.25
	impuesto := calcularImpuesto(salario)
	fmt.Printf("%.2f\n", impuesto)

}
