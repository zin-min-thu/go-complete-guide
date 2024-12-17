package main

import "fmt"

type customTransformFun func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	// use function as parameter, just use function name without ()
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

}

func transformNumbers(numbers *[]int, transform customTransformFun) []int {
	results := []int{}
	for _, val := range *numbers {
		results = append(results, transform(val))
	}

	return results
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
