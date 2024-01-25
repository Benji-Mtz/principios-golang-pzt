package main

import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

// Output personalizados de objetos
func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y la marca es: %s", myPc.ram, myPc.disk, myPc.brand)
}

func main() {
	myPC := pc{ram: 16, brand: "msi", disk: 1000}

	fmt.Println(myPC)
}
