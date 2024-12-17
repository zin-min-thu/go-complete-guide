package main

import (
	"fmt"
)

type customTransformFun func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}

	// use function as parameter, just use function name without ()
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	// get function will return function name eg(double, triple)
	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	moretransformedNumbers := transformNumbers(&moreNumbers, transformerFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(moretransformedNumbers)

}

func transformNumbers(numbers *[]int, transform customTransformFun) []int {
	results := []int{}
	for _, val := range *numbers {
		results = append(results, transform(val))
	}

	return results
}

// function return
func getTransformerFunction(numbers *[]int) customTransformFun {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
