package main

import (
	"errors"
	"fmt"
)

type salaryError struct{}

func CalcularSalario(horas int, valorHora float64) (float64, error) {
	if horas < 0 || horas < 80 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	} else {
		salario := valorHora * float64(horas)
		if salario > 150000 {
			salario -= salario * 0.10
		}
		return salario, nil
	}
}

func CalcularMedioAguinaldo(mejorSalario float64, mesesTrabajados int) (float64, error) {
	if mejorSalario < 0 || mesesTrabajados < 0 {
		return 0, errors.New("No puede ingresar numeros negativos")
	} else {
		return mejorSalario / 12 * float64(mesesTrabajados), nil
	}
}

func main() {
	salary, err := CalcularSalario(-1, 130.50)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El salario mensual es %.2f\n", salary)
		aguinaldo, err2 := CalcularMedioAguinaldo(salary, 6)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Printf("El aguinaldo es %.2f\n", aguinaldo)
		}

	}

}
