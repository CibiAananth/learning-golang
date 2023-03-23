package main

import "fmt"

const SomePublicVariable = "SomePublicValue" // public because the first letter is capital

// cant_use_walrus_here := "only inside functions"

func main() {
	var username string = "Cibi Aananth"
	fmt.Println(username)
	fmt.Printf("Type is %T \n", username)

	var truthy bool = true
	fmt.Println(truthy)
	fmt.Printf("Type is %T \n", truthy)

	var number int = 550000
	fmt.Println(number)
	fmt.Printf("Type is %T \n", number)

	var float float64 = 30231.3454232
	fmt.Println(float)
	fmt.Printf("Type is %T \n", float)

	var user_email = "cibiaananth@gmail.com"
	fmt.Println(user_email)
	fmt.Printf("Type is %T \n", user_email)

	walrus := "walrus"
	fmt.Println(walrus)
	fmt.Printf("Type is %T \n", walrus)

	fmt.Println(SomePublicVariable)
	fmt.Printf("Type is %T \n", SomePublicVariable)
}
