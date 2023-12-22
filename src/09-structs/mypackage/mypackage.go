package mypackage

import "fmt"

type CarPublic struct {
	Brand string
	Year  int
	model string
}

type carPrivate struct {
	brand string
	year  int
}

func PrintMessage() {
	fmt.Println("Hola desde el paquete")
}
