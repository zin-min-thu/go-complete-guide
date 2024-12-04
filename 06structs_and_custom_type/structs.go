package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// ... using receiver methods with struct
func (u *user) outputUserDetails() {
	fmt.Printf("First Name: %v\nLast Name: %v\nDOB: %v\n", u.firstName, u.lastName, u.birthDate)
}

// ... mutation methods with pointer receiver
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// ... constructor function and validation
func newUser(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name and birthdate are required.")
	}
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserData("Enter first name: ")
	userLastName := getUserData("Enter last name: ")
	userBirthdate := getUserData("Enter birthdate (MM/DD/YYY): ")

	var appUser *user

	// appUser = user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthDate: userBirthdate,
	// 	createdAt: time.Now(),
	// }

	appUser, err := newUser(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
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
	// fmt.Scan(&value*user)
	// Scanln to cover when we press enter empty string
	fmt.Scanln(&value)
	return value
}
