package main

import "fmt"

func main() {

	fmt.Println("Hello World")

	var count = 0

	isOdd := false

	pointer := &count

	fmt.Println(count, isOdd)

	fmt.Println("Pointer of count", pointer)
	fmt.Println("Pointer value of count", *pointer)

	type centimeter int

	type feet int

	var height1 centimeter = 175

	var height2 feet = 5

	/**
	the idea here is they are both different
	so if assign them to each other it will throw error

	e.g. height1 = height2 , correct way height1 = centimeter(height2)

	**/

	fmt.Println(height1, height2)

}
