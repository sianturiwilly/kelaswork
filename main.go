package main

import (
	"errors"
	"fmt"
)

func main() {
	var variableName1 string = "hello world"
	variableName2 := "hello world"

	fmt.Printf(variableName1)
	fmt.Printf(variableName2)

	// Primitive type: boolean, int, float, string

	// Boolean
	boolVar := false

	fmt.Printf("Type: %T Value: %v\n", boolVar, boolVar)

	// Integer
	intVar := int(5)
	intVar1 := int32(6)
	intVar2 := int64(7)

	fmt.Printf("Type: %T Value: %v\n", intVar, intVar)
	fmt.Printf("Type: %T Value: %v\n", intVar1, intVar1)
	fmt.Printf("Type: %T Value: %v\n", intVar2, intVar2)

	// Float
	floatVar1 := float32(3.2)
	floatVar2 := float64(3.2)

	fmt.Printf("Type: %T Value: %v\n", floatVar1, floatVar1)
	fmt.Printf("Type: %T Value: %v\n", floatVar2, floatVar2)

	// Bytes
	bytesVar := byte(255)
	fmt.Printf("Type: %T Value: %v\n", bytesVar, bytesVar)

	bytesVar2 := []byte("hello world")
	fmt.Printf("Type: %T Value: %v\n", bytesVar2, bytesVar2)

	// Rune
	runeVar := '😃'
	fmt.Printf("Type: %T Value: %v\n", runeVar, runeVar)

	// Complex
	complexVar := (-7 + 3i)
	fmt.Printf("Type: %T Value: %v\n", complexVar, complexVar)

	// Error
	errVar := errors.New("error detected")
	fmt.Printf("Type: %T Value: %v\n", errVar, errVar)

	// Interface
	var myInterfaceVar interface{}

	myInterfaceVar = 5
	fmt.Printf("Type: %T Value: %v\n", myInterfaceVar, myInterfaceVar)
	myInterfaceVar = "hello world"
	fmt.Printf("Type: %T Value: %v\n", myInterfaceVar, myInterfaceVar)
}

type MethodList interface {
	MyFunction()
	MyFunction2(int) int
}
