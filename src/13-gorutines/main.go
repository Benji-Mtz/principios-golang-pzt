package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	// Agrupamos un conjunto de Go routines para irlas liberando
	var wg sync.WaitGroup

	fmt.Println("Hello ")
	wg.Add(1)

	// con go se hara concurrente
	go say("World!", &wg)

	// Esperara hasta el fina es decir el defer para terminar el main
	wg.Wait()

	// time.Sleep(time.Second * 1)

}
