package main

import "fmt"

func main() {

	var days []string = []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}

	fmt.Println("Type 1")
	for d := 0; d < len(days); d++ {
		fmt.Println("type1 day ->", days[d])
	}

	fmt.Println("Type 2")
	for d, day := range days {
		fmt.Println("type2 day ->", d, day)
	}

	fmt.Println("Type 3, while loop")
	var arbitraryValue int = 1
	for arbitraryValue < 10 {
		if arbitraryValue == 2 {
			goto label1
		}
		if arbitraryValue == 4 {
			arbitraryValue++
			continue
		}
		if arbitraryValue == 7 {
			break
		}
		fmt.Println("Number", arbitraryValue)
		arbitraryValue++

	}

label1:
	fmt.Println("In label")

}
