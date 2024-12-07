package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/zin-min-thu/workingwithpackages/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		fmt.Println("--------------")
		panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")

	fmt.Println("Contact us 24/7 : ", randomdata.PhoneNumber())

	for {
		preesentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// using switch statement
		switch choice {
		case 1:
			fmt.Println("Your balance is ", accountBalance)
		case 2:
			fmt.Print("Enter your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount! Must be greater than 0.")
				// return
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Enter withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount! Must be greater than 0.")
				return
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount! You can't withdraw more than you have.")
				return
			}

			accountBalance -= withdrawAmount
			fmt.Println("Withdraw successful! New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			// break is default keword of switch statement if want to exit loop use return
			// break
			fmt.Println("Thank for choosing our bank.")
			return
		}

	}

	// fmt.Println("Thank for choosing our bank.")
}
