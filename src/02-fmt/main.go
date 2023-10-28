package main

import "fmt"

func main() {
	hello := "Hello"
	world := "World"

	// Println: imprime texto con salto de linea
	fmt.Println(hello, world)

	// Printf: inyecta variables en el texto
	nombre := "Benji"
	cursos := 20
	fmt.Printf("%s tiene más de %d cursos en Udemy\n", nombre, cursos)
	fmt.Printf("%v tiene más de %v cursos en Udemy\n", nombre, cursos)

	// Sprintf: crea cadenas para almacenarse en variables
	message := fmt.Sprintf("%s tiene más de %d cursos en Udemy", nombre, cursos)
	fmt.Println(message)

	// Tipo de datos
	fmt.Printf("hello: %T\n", hello)
	fmt.Printf("cursos: %T\n", cursos)

}
