package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])

	//panic: runtime error: index out of range [2] with length 2
	// prices[1] = 9.99
	// prices[2] = 11.99
	// fmt.Println(prices)

	prices[1] = 9.99
	prices = append(prices, 5.99)
	fmt.Println(prices)
}

/*
func main() {
	var productNames [4]string = [4]string{"ABook"}
	prices := [4]float64{22.3, 33.4, 44.5, 55.6}

	productNames[2] = "Apple"

	fmt.Println(prices)
	fmt.Println(productNames)
	fmt.Println(prices[2])

	// sample slices
	// featurePrices := prices[1:3]
	featurePrices := prices[1:]
	featurePrices[0] = 199.99
	highlightedPrices := featurePrices[:1]
	// featurePrices := prices[:3]
	fmt.Println(prices)
	fmt.Println(featurePrices)
	fmt.Println(highlightedPrices)
	// fmt.Println(len(featurePrices), cap(featurePrices)) // outpub 3 3
	fmt.Println(len(highlightedPrices), cap(highlightedPrices)) // outpub 1 3

	highlightedPrices = highlightedPrices[:3] // we can slice only for right side(eg: :3)
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices)) // outpub 3 3
}
*/
