package main

import (
	"fmt"
	"log"
	"os"
)

const path string = "./info.csv"

func main() {
	id := 00
	precio := 10
	cantidad := 1
	csv := "ID, Precio, Cantidad\n"

	for i := 0; i < 10; i++ {
		line := fmt.Sprintf("%d,%d,%d\n", id, precio, cantidad)
		csv += line
		id++
		precio++
		cantidad++
	}

	bytes := []byte(csv)
	err := os.WriteFile(path, bytes, os.FileMode(0644))
	if err != nil {
		log.Fatal(err)
	}
}
