package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func main() {

	// En los punteros el valor de memoria es & pero el valor perse es *

	a := 50
	b := &a

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(*b)

	// Como a = 50 y *b = 50 cualquiera sirve para editar el valor de a
	*b = 100
	fmt.Println(a)

	// var myPc pc
	myPc := pc{ram: 16, disk: 256, brand: "mac"}
	fmt.Println(myPc)

	myPc.ping()

	fmt.Println(myPc)
	myPc.duplicateRAM()
	fmt.Println(myPc)
	myPc.duplicateRAM()
	fmt.Println(myPc)

}
