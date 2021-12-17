package main

import "fmt"

func main() {
	word := "palabra"
	len := len(word)
	fmt.Println(len)
	for i := range word {
		fmt.Printf("%c\n", word[i])
	}
}
