package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const fileName string = "./ejercicio1/info.csv"

func main() {
	file, err := os.Open(fileName)

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
			fmt.Printf("%s\t\t%s\t%s\n", line[0], line[1], line[2])
		}

	}
}
