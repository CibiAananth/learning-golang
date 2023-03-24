package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	languageMap := make(map[int]string)
	languageMap[0] = "JS"
	languageMap[1] = "Ruby"
	languageMap[2] = "GoLang"
	languageMap[4] = "Python"
	languageMap[3] = "Rust"
	fmt.Println("The map is", languageMap)

	delete(languageMap, 1)
	fmt.Println("The map after delete is", languageMap)

	for key, value := range languageMap {
		fmt.Printf("%v -> %v \n", key, value)
	}
}
