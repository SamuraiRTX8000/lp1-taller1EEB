package main

import (
	"fmt"
	"sync"
	"time"
)

// Objetivo: Lanzar varias goroutines que imprimen mensajes y esperar a que todas terminen.
// TODO: Completa los pasos marcados con TODO para entender goroutines y WaitGroup.

func worker(id int, veces int, wg *sync.WaitGroup) {

	for i := 1; i <= veces; i++ {
		fmt.Printf("[worker %d] hola %d\n", id, i)

		//time.sleep
		time.Sleep(300 * time.Millisecond)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	// numGoroutines
	// veces

	numGoroutines := 10
	veces := 3

	for id := 1; id <= numGoroutines; id++ {
		wg.Add(1)
		go worker(id, veces, &wg)

	}
	wg.Wait()

	fmt.Println("Listo: todas las goroutines terminaron.")
}
