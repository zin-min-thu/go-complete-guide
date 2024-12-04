package user

import (
	"errors"
	"fmt"
	"time"
)

// User Capitalize public can call outsite of the other files with package name
// user small private can use within a file
/* also function name and variable name are public(Capital letter) and private(small letter), we can support*/

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// ... using receiver methods with struct
func (u *User) OutputUserDetails() {
	fmt.Printf("First Name: %v\nLast Name: %v\nDOB: %v\n", u.firstName, u.lastName, u.birthDate)
}

// ... mutation methods with pointer receiver
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// ... constructor function and validation
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}
