package main

import (
	"fmt"
)

type Fecha struct {
	Dia, Mes, Anio int
}

type Estudiante struct {
	Nombre       string
	Apellido     string
	DNI          string
	FechaIngreso Fecha
}

func (e Estudiante) Detalle() {
	fmt.Printf("%+v\n", e)
}

func main() {
	e1 := Estudiante{"Rodrigo", "Souza", "5006171-7", Fecha{13, 12, 2021}}
	e1.Detalle()
}
