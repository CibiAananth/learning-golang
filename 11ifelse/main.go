package main

import "fmt"

func main() {
	fmt.Println("If else")

	const print_if bool = true
	if print_if {
		fmt.Println("Inside 'if' block")
	} else {
		fmt.Println("Inside 'else' block")
	}

	if print_else := true; !print_else {
		fmt.Println("Inside 'if' block")
	} else {
		fmt.Println("Inside 'else' block")
	}
}
