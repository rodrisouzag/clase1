package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

func print() {
	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("failed to open")

	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	for i, each_ln := range text {
		line := strings.Split(each_ln, ",")
		if i == 0 {
			fmt.Printf("%s\t%s\t%s\n", line[0], line[1], line[2])
		} else {
			fmt.Printf("%s\t%s\t\t%s\n", line[0], line[1], line[2])
		}

	}
}
