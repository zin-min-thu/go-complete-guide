package main

import (
	"errors"
	"fmt"
	"os"
)

func storeResultFile(ebt, profit, ratio float64) {
	// ebtText := fmt.Sprint()
	// outputText := fmt.Sprint("EBT: ", ebt, "\nProfit ", profit, "\nRatio ", ratio)
	outputText := fmt.Sprintf("EBT: %.1f\nProfit: %.2f\nRatio: %.3f\n", ebt, profit, ratio)

	os.WriteFile("profit.txt", []byte(outputText), 0644)
}

func main() {

	// var revenue, expenses float64
	// var taxRate float64

	// fmt.Print("Enter revenue: ")
	// fmt.Scan(&revenue)
	revenue, err := getUserInput("Enter revenue: ")

	if err != nil {
		fmt.Println(err)
		return
		// panic(err)
	}

	// fmt.Print("Enter expenses: ")
	// fmt.Scan(&expenses)
	expenses, err := getUserInput("Enter expenses: ")

	if err != nil {
		panic(err)
	}

	// fmt.Print("Enter taxRate: ")
	// fmt.Scan(&taxRate)
	taxRate, err := getUserInput("Enter taxRate: ")

	if err != nil {
		panic(err)
	}

	ebt, profit, ratio := calculateFinancial(revenue, expenses, taxRate)

	// fmt.Println("EBT : ", ebt)
	// fmt.Println("Profit : ", profit)
	// fmt.Println("Ratio: ", ratio)

	storeResultFile(ebt, profit, ratio)

	fmt.Printf("EBT : %v\n", ebt)
	fmt.Printf("Profit : %v\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)

}

func getUserInput(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return userInput, errors.New("invalid input number, you enter less than zero")
	}

	return userInput, nil
}

func calculateFinancial(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}
