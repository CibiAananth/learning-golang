package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Type conversion")
	var reader = bufio.NewReader(os.Stdin)
	rating, _ := reader.ReadString('\n')

	converted, err := strconv.ParseInt(strings.TrimSpace(rating), 0, 64)

	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("rating is", converted+1)
	}
}
