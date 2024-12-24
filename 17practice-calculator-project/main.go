package main

import (
	"fmt"

	"github.com/zin-min-thu/practicecalculatorproject/filemanager"
	"github.com/zin-min-thu/practicecalculatorproject/prices"
)

func main() {
	// prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	// results := make(map[float64][]float64)

	for _, taxRate := range taxRates {

		fm := filemanager.New("pricesss.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))

		// cmdm := cmdmanager.New()

		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could no process job")
			fmt.Println(err)
		}
	}

}
