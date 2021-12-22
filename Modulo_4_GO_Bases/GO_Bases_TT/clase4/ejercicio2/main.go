package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func getLegajo() *int {
	nro := new(int)
	random := rand.Intn(100)
	nro = &random
	nro = nil
	return nro
}

type Cliente struct {
	Legajo    int
	Nombre    string
	DNI       string
	Telefono  int
	Domicilio string
}

func openFile(filename string) *os.File {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()

	file, err := os.Open(filename)

	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	} else {
		return file
	}
}

func validarDatos(nombre string, dni string, telefono int, domicilio string) (bool, error) {
	if nombre == "" || dni == "" || telefono == 0 || domicilio == "" {
		return false, errors.New("ha ingresado un dato vacio")
	} else {
		return true, nil
	}
}

func main() {

	defer func() {
		fmt.Println("fin de la ejecucion")
		fmt.Println("se detectaron varios errores en tiempo de ejecucion")
		fmt.Println("no han quedado archivos abiertos")
	}()

	nroLegajo := getLegajo()
	if nroLegajo != nil {
		openFile("./customers.txt")
		_, err := validarDatos("Nombre", "DNI", 0, "Domicilio")
		if err != nil {

		}
	} else {
		panic("el numero de legajo no puede ser nil")
	}

}
