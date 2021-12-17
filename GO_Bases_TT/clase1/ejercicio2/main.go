package main

import (
	"fmt"
)

var price, discount float64 = 1200.0, 25.0

func main() {
	fmt.Printf("El precio con descuento es %.2f\n", ((100-discount)/100)*price)
}
