package main

import "fmt"

func main() {

	// bool

	// string

	// int int8 int16 int32 int64

	// uint uint8 uint16 uint32 uint64 uintptr

	// byte // alias for uint8 because 1 byte = 8 bits (binary)

	// rune // alias for int32
	// represents a Unicode code point or a character

	// float32 float64
	// complex64 complex128

	var text1 string = "First way of declaring variable"
	text2 := "Another way of declaring variable" // short assignment operators

	fmt.Println(text1)
	fmt.Println(text2)

	firstName, lastName := "Wai Kit", "Ho"

	fmt.Println("My name is", firstName, lastName)

	const constVariable = "constant var" // does not have short assignment operators

	age1 := 18
	if age1 >= 18 {
		fmt.Println("Carti")
	} else {
		fmt.Println("Cocomelon")
	}

	fmt.Println(mergeName(firstName, lastName))

}

func mergeName(x string, y string) string {
	return x + y
}
