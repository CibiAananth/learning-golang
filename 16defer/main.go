package main

import "fmt"

func main() {
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("defer")
	printNumbers()
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
