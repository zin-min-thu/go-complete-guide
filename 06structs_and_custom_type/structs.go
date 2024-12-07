package main

import (
	"fmt"

	"github.com/zin-min-thu/structandcustomtypes/user"
)

func main() {
	userFirstName := getUserData("Enter first name: ")
	userLastName := getUserData("Enter last name: ")
	userBirthdate := getUserData("Enter birthdate (MM/DD/YYY): ")

	var appUser *user.User

	// appUser = user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthDate: userBirthdate,
	// 	createdAt: time.Now(),
	// }

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("admin@gmail.com", "123")
	// admin.User.OutputUserDetails()
	// admin.User.ClearUserName()
	// admin.User.OutputUserDetails()
	// direct accept cuz using anonymous struct User type
	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

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
