package main

import (
	"fmt"
)

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
var cont int = 0

func main() {
	fmt.Printf("Benjamin tiene %d años\n", employees["Benjamin"])
	for i := range employees {
		if employees[i] > 21 {
			cont++
		}
	}
	fmt.Printf("Hay %d mayores\n", cont)
	employees["Federico"] = 25
	fmt.Println(employees)
	delete(employees, "Pedro")
	fmt.Println(employees)
}
