package main

import (
	"fmt"
)

type customTransformFun func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	double := createTransformer(2)
	triple := createTransformer(3)
	// using anonymous function
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	results := []int{}
	for _, val := range *numbers {
		results = append(results, transform(val))
	}

	return results
}

// closures
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
