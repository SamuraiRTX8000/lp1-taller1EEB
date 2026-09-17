package main

import (
	"fmt"
	"sync"
	"time"
)

// Objetivo: Simular "futuros" en Go usando canales. Una función lanza trabajo asíncrono
// y retorna un canal de solo lectura con el resultado futuro.
// TODO: completa las funciones y experimenta con varios futuros a la vez.
// Asincronico 1
func asyncCuadrado(x int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)

		time.Sleep(500 * time.Millisecond)

		ch <- x * x
	}()
	return ch
}

// Asincronico Fusion de canales
func fanIn(canales ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(canales))

	for _, canal := range canales {
		go func(c <-chan int) {
			defer wg.Done()

			for valor := range c {
				out <- valor
			}
		}(canal)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {

	f1 := asyncCuadrado(2)
	f2 := asyncCuadrado(3)
	f3 := asyncCuadrado(4)

	//Para fan in los siguientes canales

	f4 := asyncCuadrado(5)
	f5 := asyncCuadrado(6)
	f6 := asyncCuadrado(7)

	//secuencialmente
	l1 := <-f1
	l2 := <-f2
	l3 := <-f3

	fmt.Printf("Resultados: %d, %d, %d\n", l1, l2, l3)

	canal := fanIn(f4, f5, f6)

	i := 4
	for valor := range canal {

		fmt.Printf("Resultado fan-in Chanel: %d %d\n", i, valor)
		i = i + 1 //solo se declara con := si la funcion es la primer vez que se crea si se
		//sobre escribe se usa = para asignar un nuevo valor a la variable ya creada
	}
}
