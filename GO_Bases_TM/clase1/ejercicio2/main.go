package main

import (
	"fmt"
)

var temp float64 = 22.0
var humidity, press int = 69, 1017

func main() {
	fmt.Println("La temperatura es:", temp)
	fmt.Println("La humedad es:", humidity)
	fmt.Println("La presion atmosferica es:", press)
}
