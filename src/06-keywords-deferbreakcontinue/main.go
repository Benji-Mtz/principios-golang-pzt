package main

import "fmt"

func main() {
	// Defer
	defer fmt.Println("Hola ")
	fmt.Println("Mundo!")

	// Break & Continue
	for i := 0; i <= 10; i++ {
		fmt.Println(i)

		// Continue: Para hacer algo y que siga el proceso
		if i == 2 {
			fmt.Println("es 2")
			continue
		}

		// Break: Cuando sucede algo y requieres detener el proceso
		if i == 8 {
			fmt.Println("Llegamos a 8, saliendo...")
			break
		}
	}
}
