package main

import "fmt"

func main() {
	fmt.Println("Array")

	var languages [3]string
	languages[0] = "Python"
	languages[1] = "JS"
	languages[2] = "Rust"

	fmt.Println("Languages", languages)
	fmt.Println("Length", len(languages))

	var frameworks = [2]string{"Go", "JavaScript"}

	// To create an array size based on number of elements
	// var frameworks = [...]string{"Go", "JavaScript"}

	// This will give error since array are not constants and can be mutated
	// const frameworks = [2]string{"Go", "JavaScript"}

	fmt.Println("Frameworks", frameworks)
	fmt.Println("Length", len(frameworks))

	const (
		A int = 1
		B int = 2
	)

	arrayFromValues := [...]int{A, B} // length of array will be the number of elements
	fmt.Println("ArrayFromValues", arrayFromValues)
	fmt.Println("Length", len(arrayFromValues))

	slicedArray := arrayFromValues[:1] // slice [start:end], blank for start and end of array
	fmt.Println("SlicedArray", slicedArray)
	fmt.Println("Length", len(slicedArray))
}
