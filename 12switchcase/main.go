package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch case")

	var randomNumber = rand.Intn(5) + 1
	switch randomNumber {
	case 1:
		fmt.Println("The number is 1")
		fallthrough
	case 2:
		fmt.Println("The number is 2")
		fallthrough
	case 3:
		fmt.Println("The number is 3")
		fallthrough
	case 4:
		fmt.Println("The number is 4")
		fallthrough
	case 5:
		fmt.Println("The number is 5")
	}
}
