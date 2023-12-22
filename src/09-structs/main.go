package main

// Is posible that we have to set a environment variable like: go env -w GO111MODULE=off
import (
	"fmt"

	mialias "github.com/Benji-Mtz/golang-ptz/src/09-structs/mypackage"
)

func main() {
	// Primer método
	var myCar mialias.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2023
	fmt.Println(myCar)

	mialias.PrintMessage()

	/*
		var otherCar car
		otherCar.brand = "Ferrari"
		fmt.Println(otherCar)
	*/
}
