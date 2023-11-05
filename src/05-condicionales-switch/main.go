package main

import "fmt"

func main() {
	// Mis variables
	valor1, valor2 := 1, 2
	// modulo := 4 % 2

	if valor1 == 1 {
		fmt.Println("es 1")
	} else {
		fmt.Println("No es 1")
	}

	// And
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("condición verdadera")
	} else {
		fmt.Println("La condición no se cumple")
	}

	// Or
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es verdad OR")
	}

	// Switch con condición añadida

	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// Switch sin condición inicial
	value := 200
	switch {
	case value < 0:
		fmt.Println("valor menor a cero")
	case value > 100:
		fmt.Println("valor mayor a 100")
	default:
		fmt.Println("El valor esta entre 1 y 99")
	}
}
