package main

type fuguras2D interface {
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

func calcular(f fuguras2D)

func main() {
	myCuadrado := cuadrado{base: 5}
	myRectangulo := rectangulo{base: 2, altura: 4}
}
