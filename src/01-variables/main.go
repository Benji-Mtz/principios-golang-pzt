package main

import "fmt"

func main() {
	// Declaración de constantes
	const pi float64 = 3.1416
	const pi2 = 3.15

	fmt.Println("PI", pi)
	fmt.Println("PI2", pi2)

	// Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero Values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area del cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado + baseCuadrado
	fmt.Println("Area del cuadrado", areaCuadrado)

}
