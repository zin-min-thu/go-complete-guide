package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	accountBalance, err := getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		fmt.Println("--------------")
		panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			// break is default keword of switch statement if want to exit loop use return
			// break
			fmt.Println("Thank for choosing our bank.")
			return
		}

		// using if statement
		// if choice == 1 {
		// 	fmt.Println("Your balance is ", accountBalance)
		// } else if choice == 2 {
		// 	fmt.Print("Enter your deposit: ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)

		// 	if depositAmount <= 0 {
		// 		fmt.Println("Invalid amount! Must be greater than 0.")
		// 		// return
		// 		continue
		// 	}

		// 	accountBalance += depositAmount
		// 	fmt.Println("Balance updated! New amount: ", accountBalance)
		// } else if choice == 3 {
		// 	fmt.Print("Enter withdraw amount: ")
		// 	var withdrawAmount float64
		// 	fmt.Scan(&withdrawAmount)

		// 	if withdrawAmount <= 0 {
		// 		fmt.Println("Invalid amount! Must be greater than 0.")
		// 		continue
		// 	}

		// 	if withdrawAmount > accountBalance {
		// 		fmt.Println("Invalid amount! You can't withdraw more than you have.")
		// 		continue
		// 	}

		// 	accountBalance -= withdrawAmount
		// 	fmt.Println("Withdraw successful! New amount: ", accountBalance)
		// } else {
		// 	fmt.Println("Goodbye!")
		// 	// return can only be infinit loop to exit
		// 	//return
		// 	// break can work out of the infinity loop
		// 	break
		// }

	}

	// fmt.Println("Thank for choosing our bank.")
}
