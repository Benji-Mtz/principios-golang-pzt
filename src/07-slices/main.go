package main

import (
	"fmt"
	"strings"
)

func esPalindromo(text string) {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	text = strings.ToLower(text)
	textReverse = strings.ToLower(textReverse)

	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func main() {
	// Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	// array[2] = 2
	// array[3] = 2
	fmt.Println(array, len(array), cap(array))

	// Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])
	fmt.Println(slice[:])

	// Append a element
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append new list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	// slice string
	sliceString := []string{"Hola", "que", "hace"}

	for i, valor := range sliceString {
		fmt.Println(i, valor)
	}

	// funcion palindromo
	esPalindromo("Ama")
}
