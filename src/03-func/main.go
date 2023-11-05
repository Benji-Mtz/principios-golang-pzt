package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArguments(a, b int, c string) {
	fmt.Println(a, b, c)
}

func squareValue(a int) int {
	return a * a
}

func doubleReturn(a int) (b, c int) {
	return a, a * 2
}

func main() {

	normalFunction("Hola Mundo!")
	tripleArguments(1, 2, "Benji")

	cinco := squareValue(5)
	fmt.Println(cinco)

	valOne, valTwo := doubleReturn(10)
	fmt.Printf("val1: %d y val2: %d \n", valOne, valTwo)
}
