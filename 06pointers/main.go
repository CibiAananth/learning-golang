package main

import "fmt"

func main() {
	fmt.Println("My pointer")

	var ptr *int
	fmt.Println("Pointer def value is", ptr)

	var number int = 2
	numberPtr := &number

	fmt.Println("The pointer value is", numberPtr)
	fmt.Println("The value in address is", *numberPtr)

	*numberPtr = *numberPtr * 2
	fmt.Println("The updated value is", number)

}
