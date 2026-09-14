package main

import (
	"fmt"
	"sync"
	"time"
)

// Objetivo: Simular tareas que toman tiempo con time.Sleep y comparar
// ejecución secuencial vs concurrente midiendo la duración total.
// TODO: completa las secciones marcadas para observar la mejora.

func tarea(id int, dur time.Duration) {
	fmt.Printf("[tarea %d] iniciando, dur=%v\n", id, dur)
	time.Sleep(dur)
	fmt.Printf("[tarea %d] finalizada\n", id)
}

func secuencial(durs []time.Duration) time.Duration {
	inicio := time.Now()

	for i, d := range durs {
		tarea(i, d)

	}
	return time.Since(inicio)
}

func concurrente(durs []time.Duration) time.Duration {
	inicio := time.Now()
	var wg sync.WaitGroup

	for i, d := range durs {
		wg.Add(1)
		go func(id int, duration time.Duration) {
			defer wg.Done()
			tarea(id, duration)
		}(i, d)
	}
	wg.Wait()
	return time.Since(inicio)
}

func main() {
	// TODO: experimenta con diferentes duraciones
	durs := []time.Duration{1 * time.Second, 3 * time.Second, 6 * time.Second}
	//modificaciones para el tiempo fueron:1,3,6
	// d1 :=
	d1 := secuencial(durs)
	fmt.Println("Duración SEC:", d1)

	// d2 :=
	d2 := concurrente(durs)
	fmt.Println("Duración CONC:", d2)

	fmt.Println("Nota: la ejecución concurrente debería ser ~max(durs). Cambia valores y observa.")
}
