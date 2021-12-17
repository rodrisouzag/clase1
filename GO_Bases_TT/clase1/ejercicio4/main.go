package main

import (
	"fmt"
)

var index int = 12

func main() {
	month := map[int]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	fmt.Println(month[index])
}
