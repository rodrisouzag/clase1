package main

var age int = 23
var employed bool = true
var seniority int = 2
var salary float64 = 1

func main() {
	if age > 22 && employed && seniority > 1 && salary > 100000.0 {
		println("Prestamo aprobado sin intereses")
	} else if age > 22 && employed && seniority > 1 && salary < 100000.0 {
		println("Prestamo aprobado con intereses")
	} else {
		println("Prestamo rechazado")
	}
}
