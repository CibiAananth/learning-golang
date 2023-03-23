package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Prompt")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter reader")

	// comma ok || err ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Input is", input)
}
