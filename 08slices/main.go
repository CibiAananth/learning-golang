package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in GoLang")

	var languages = [1]string{"JS"} // array
	fmt.Println("Languages 1", languages)

	var languages2 = []string{} // slices
	fmt.Println("Languages 2", languages2)

	var languages3 = []string{"GoLang", "Python", "Rust", "C", "C++"} // slices
	fmt.Println("Languages 3", languages3)

	languages3 = append(languages3, "JS")
	fmt.Println("Languages 3 after append", languages3)

	sort.Strings(languages3)
	fmt.Println("Languages 3 post sort", languages3)

	languages4 := languages3[:2]
	fmt.Println("Languages 4", languages4)

	const indexToRemove int = 1
	languages5 := append(languages3[:indexToRemove], languages3[indexToRemove+1:]...)
	fmt.Println("Languages 5", languages5)

	// languages4 = append(languages3[:]) // this is similar to languages4 = languages3

	numberSlice := make([]int, 4) // initialized with default value 0
	numberSlice[0] = 2
	numberSlice[1] = 1
	numberSlice[2] = 5
	numberSlice[3] = 3
	fmt.Println("numberSlice", numberSlice)

	numberSlice = append(numberSlice, 12, 7) // memory allocation happens again
	fmt.Println("numberSlice post append", numberSlice)

	sort.Ints(numberSlice)
	fmt.Println("numberSlice post sort", numberSlice)

}
