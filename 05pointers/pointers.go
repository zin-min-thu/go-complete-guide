package main

import "fmt"

func main() {
	age := 32 // regular variable

	// we can declare
	// var agePointer *int
	agePointer := &age

	fmt.Println("Age: ", age)
	fmt.Println("Age Pointer: ", agePointer)
	// if you want to get value behind the pointer
	fmt.Println("Age Pointer Value: ", *agePointer)

	// if we use age in different function, go will copy of the value and use two momery address
	// adultYears := getAdultYears(agePointer)
	// fmt.Println(adultYears)

	editAgeToAdultYears(agePointer)
	fmt.Println("Dereference Age: ", age)

	// testing default empty variables pointer
	/*
		var intVal int
		var floatVal float64
		var stringVal string

		intPointer := &intVal
		floatPointer := &floatVal
		stringPointer := &stringVal

		fmt.Println("Int pointer: ", &intPointer)
		fmt.Println("Int pointer value: ", *intPointer)
		fmt.Println("Float pointer: ", &floatPointer)
		fmt.Println("Float pointer value: ", *floatPointer)
		fmt.Println("String pointer: ", &stringPointer)
		fmt.Println("String pointer value: ", *stringPointer)
	*/

}

// func getAdultYears(age *int) int {
// 	return *age - 18
// }

// deference pointer
func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
