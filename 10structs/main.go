package main

import "fmt"

func main() {
	fmt.Println("Structs")

	type User struct {
		ID    string
		Name  string
		Email string
	}

	var user_cibi_1 User = User{"1", "Cibi Aananth", "xyz@gmail.com"}
	fmt.Printf("user cibi 1 %+v\n", user_cibi_1)

	var user_cibi2 = new(User)
	user_cibi2.Email = "xyz@gmail.com"
	user_cibi2.ID = "2"
	user_cibi2.Name = "Cibi Aananth"
	fmt.Printf("user cibi 2 %+v\n", *user_cibi2)

	var user_cibi3 = &User{"3", "Cibi Aananth", "xyz@gmail.com"}
	fmt.Printf("user cibi 2 %+v\n", *user_cibi3)

}
