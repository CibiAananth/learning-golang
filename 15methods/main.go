package main

import "fmt"

type User struct {
	Name     string
	Email    string
	Age      int
	IsActive bool
}

func (u User) GetStatus() bool {
	return u.IsActive
}

func (u User) UpdateEmail(email string) {
	u.Email = email
	fmt.Println("new email is", u.Email)
}

func (u *User) UpdateEmailInOriginal(email string) {
	u.Email = email
	fmt.Println("new email is", u.Email)
}

func main() {
	fmt.Println("methods")

	var user1 User = User{"Cibi", "cibi@test.com", 29, true}
	fmt.Printf("user1 %+v\n", user1)
	fmt.Println("status", user1.GetStatus())
	user1.UpdateEmail("cibiw@test.com")
	fmt.Println("email from main", user1.Email)
	user1.UpdateEmailInOriginal("cibinew@test.com")
	fmt.Println("email from main", user1.Email)
}
