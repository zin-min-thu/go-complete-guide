package main

import "fmt"

func main() {
	resultInt := add(1, 2)
	fmt.Println(resultInt)

	resultFloat := add(1.2, 2.3)
	fmt.Printf("%.2f\n", resultFloat)

	resultString := add("A", "B")
	fmt.Println(resultString)
}

// dynamic type & limitation
// func add(a, b interface{}) interface{} {
// 	aInt, aIsInt := a.(int)
// 	bInt, bIsInt := b.(int)

// 	if aIsInt && bIsInt {
// 		return aInt + bInt
// 	}

// 	aFloat, aIsFloat := a.(int)
// 	bFloat, bIsFloat := b.(int)

// 	if aIsFloat && bIsFloat {
// 		return aFloat + bFloat
// 	}

// 	aString, aIsString := a.(int)
// 	bString, bIsString := b.(int)

// 	if aIsString && bIsString {
// 		return aString + bString
// 	}

// 	return nil
// }

// working with generics
func add[T int | float64 | string](a, b T) T {
	return a + b
}
