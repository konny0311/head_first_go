package main

import "fmt"

func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}

func main() {
	var tmp = 10
	var myFloatPointer *float64 = createPointer()
	fmt.Println(myFloatPointer)
	fmt.Println(*myFloatPointer)

	fmt.Println(&tmp)
}
