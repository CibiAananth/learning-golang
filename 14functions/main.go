package main

import "fmt"

func main() {
	fmt.Println("Function")
	myFunction()
}

func myFunction() {
	var result int = adder(1, 2)
	fmt.Println("result", result)

	proResult, message := proAdder(1, 5)
	fmt.Println("proResult", proResult)

	fmt.Println("message", message)
}

func adder(x int, y int) int {
	return x + y
}

func proAdder(values ...int) (int, string) {
	var total int = 0

	for _, val := range values {
		total += val
	}

	return total, "Hello"
}
