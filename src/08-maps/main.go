package main

import (
	"fmt"
)

func main() {
	// Maps o diccionarios

	dict := make(map[string]int)

	dict["Benji"] = 36
	dict["Michelle"] = 39

	fmt.Println(dict)

	for indice, valor := range dict {
		fmt.Printf("Indice: %s, Valor:%d\n", indice, valor)
	}

	// Encontrar un valor por su llave
	// var name string = "Pepe"
	var name string = "Benji"
	valor, isInMap := dict[name]
	fmt.Printf("%s esta en el diccionario? %t, su valor es:%d\n", name, isInMap, valor)
}
