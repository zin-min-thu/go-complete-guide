package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	// sum := sumup(numbers)

	sum := sumup(1, 10, 15, 20, -5)

	anotherSum := sumup(1, numbers...)

	fmt.Println("Sum = ", sum)
	fmt.Println("Another Sum = ", anotherSum)
}

// func sumup(numbers []int) int {
// variadic function -> startingValue = 1, numbers = [10,15,20,-5]
func sumup(startingValue int, numbers ...int) int {
	sum := 0

	fmt.Println("Starting Value: ", startingValue)
	fmt.Println(numbers)

	for _, val := range numbers {
		sum += val
	}

	return sum
}
