package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	slice := generateSlice(100)
	b, s, i := bubblesort(slice), selectionsort(slice), insertionsort(slice)
	fmt.Printf("Tiempo total %d microsegundos", <-b+<-s+<-i)

	slice2 := generateSlice(1000)
	b, s, i = bubblesort(slice2), selectionsort(slice2), insertionsort(slice2)
	fmt.Printf("Tiempo total %d microsegundos", <-b+<-s+<-i)

	slice3 := generateSlice(10000)
	b, s, i = bubblesort(slice3), selectionsort(slice3), insertionsort(slice3)
	fmt.Printf("Tiempo total %d microsegundos", <-b+<-s+<-i)
}

func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func bubblesort(items []int) <-chan int {
	r := make(chan int)
	go func() {
		inicio := time.Now().UnixMicro()
		var (
			n      = len(items)
			sorted = false
		)
		for !sorted {
			swapped := false
			for i := 0; i < n-1; i++ {
				if items[i] > items[i+1] {
					items[i+1], items[i] = items[i], items[i+1]
					swapped = true
				}
			}
			if !swapped {
				sorted = true
			}
			n = n - 1
		}
		fin := time.Now().UnixMicro()
		fmt.Println("------------------------")
		fmt.Printf("--> Bubble Sort para %d items duró %d microsegundos\n", len(items), fin-inicio)
		fmt.Println("------------------------")
		r <- int(fin - inicio)
	}()
	return r
}

func selectionsort(items []int) <-chan int {
	r := make(chan int)
	go func() {
		inicio := time.Now().UnixMicro()
		var n = len(items)
		for i := 0; i < n; i++ {
			var minIdx = i
			for j := i; j < n; j++ {
				if items[j] < items[minIdx] {
					minIdx = j
				}
			}
			items[i], items[minIdx] = items[minIdx], items[i]
		}
		fin := time.Now().UnixMicro()
		fmt.Println("------------------------")
		fmt.Printf("--> Selection Sort para %d items duró %d microsegundos\n", n, fin-inicio)
		fmt.Println("------------------------")
		r <- int(fin - inicio)
	}()
	return r
}

func insertionsort(items []int) <-chan int {
	r := make(chan int)
	go func() {
		inicio := time.Now().UnixMicro()
		var n = len(items)
		for i := 1; i < n; i++ {
			j := i
			for j > 0 {
				if items[j-1] > items[j] {
					items[j-1], items[j] = items[j], items[j-1]
				}
				j = j - 1
			}
		}
		fin := time.Now().UnixMicro()
		fmt.Println("------------------------")
		fmt.Printf("--> Insertion Sort para %d items duró %d microsegundos\n", n, fin-inicio)
		fmt.Println("------------------------")
		r <- int(fin - inicio)
	}()
	return r
}
