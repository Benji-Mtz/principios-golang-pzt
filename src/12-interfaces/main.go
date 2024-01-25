package main

import "fmt"

type figuras2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figuras2D) {
	fmt.Println("Area:", f.area())
}

func main() {
	myCuadrado := cuadrado{base: 5}
	myRectangulo := rectangulo{base: 2, altura: 4}

	// Implementando la interfaz para las areas de todas las figuras
	calcular(myCuadrado)
	calcular(myRectangulo)

	// Listas de interfaces
	myInterface := []interface{}{"Hola", 12, 1.50}
	fmt.Println(myInterface...)
}
