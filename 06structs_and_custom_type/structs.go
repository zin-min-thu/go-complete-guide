package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// using receiver methods with struct
func (u *user) outputUserDetails() {
	fmt.Printf("First Name: %v\nLast Name: %v\nDOB: %v\n", u.firstName, u.lastName, u.birthDate)
}

// mutation methods with pointer receiver
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	userFirstName := getUserData("Enter first name: ")
	userLastName := getUserData("Enter last name: ")
	userBirthdate := getUserData("Enter birthdate (MM/DD/YYY): ")

	appUser := user{}

	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()

	// sort form must be same order with user struct
	// appUser = user{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthdate,
	// 	time.Now(),
	// }

	// outputUser(appUser) // without pointer will copy appUser when call then store new memory address
	// outputUser(&appUser)
}

// using pointer
// func outputUser(u *user) {
// 	// fmt.Printf("First Name: %v\nLast Name: %v\nDOB: %v\n", firstName, lastName, birthdate)
// 	fmt.Printf("First Name: %v\nLast Name: %v\nDOB: %v\n", u.firstName, u.lastName, u.birthDate)
// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
