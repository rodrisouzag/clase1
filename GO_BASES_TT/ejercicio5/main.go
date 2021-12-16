package main

import (
	"fmt"
)

var students []string

func main() {
	students = append(students, "Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka")
	fmt.Println(students)
	// dos clases despues
	students = append(students, "Garbiela")
	fmt.Println(students)
}
