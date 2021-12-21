package main

func InsertionSort(numbers ...int) {

	for i := range numbers {
		actual := numbers[i]
		j := i
		for j > 0 && numbers[j-1] > actual {
			numbers[j] = numbers[j-1]
			j = j - 1
		}
		numbers[j] = actual

	}
}

func main() {
	//variable1 := rand.Perm(100)
	//variable2 := rand.Perm(1000)
	//variable3 := rand.Perm(10000)
}
