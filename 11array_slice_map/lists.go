package main

import "fmt"

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
	highlightedPrices := featurePrices[:1]
	// featurePrices := prices[:3]
	fmt.Println(featurePrices)
	fmt.Println(highlightedPrices)
}
