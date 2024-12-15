package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	// 1)
	hobbies := [3]string{"Sports", "Cooking", "Reading"}
	fmt.Println(hobbies)

	// 2)
	fmt.Println(hobbies[0])
	// fmt.Println(hobbies[1:])
	fmt.Println(hobbies[1:3])

	// 3)
	mainHobbies := hobbies[:2]
	// mainHobbies := hobbies[0:2]
	fmt.Println(mainHobbies)

	// 4)
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:]
	fmt.Println(mainHobbies)

	// 5)
	courseGoals := []string{"Learn Go!", "Learn all the basice"}
	fmt.Println(courseGoals)

	// 6)
	courseGoals[1] = "Learn details"
	courseGoals = append(courseGoals, "Learn second all the basic")
	fmt.Println(courseGoals)

	// 7)
	products := []Product{
		{
			"1-product", "First Product Name", 11.1,
		},
		{
			"2-product", "Second Product Name", 22.2,
		},
	}

	fmt.Println(products)

	newProduct := Product{
		"3-product",
		"Third product name",
		33.33,
	}
	products = append(products, newProduct)

	fmt.Println(products)

}
