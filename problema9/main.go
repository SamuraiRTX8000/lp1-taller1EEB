package main

import (
	"fmt"
	"sync"
	"time"
)

// Objetivo: Implementar una versión del problema de los Filósofos Comensales.
// Hay 5 filósofos y 5 tenedores (recursos). Cada filósofo necesita 2 tenedores para comer.
// Estrategia segura: imponer un **orden global** al tomar los tenedores (primero el menor ID, luego el mayor)
// para evitar deadlock. También puedes limitar concurrencia (ej. mayordomo).
// TODO: completa la lógica de toma/soltado de tenedores y bucle de pensar/comer.

type tenedor struct{ mu sync.Mutex }

func filosofo(id int, izq, der *tenedor, wg *sync.WaitGroup) {
	defer wg.Done()
	numero := rand.Intn(100)
	aBinario := aBinario(numero)
	fmt.Printf("[filósofo %d] binario: %s\n", id, aBinario)

	

	fmt.Printf("[filósofo %d] satisfecho\n", id)
}

func pensar(id int) {
	fmt.Printf("[filósofo %d] pensando...\n", id)
	time.Sleep(300 * time.Millisecond) // Simular tiempo de pensar

}

func comer(id int) {
	fmt.Printf("[filósofo %d] COMIENDO\n", id)

	time.Sleep(300 * time.Millisecond)

}

func main() {
	const n = 5
	var wg sync.WaitGroup
	wg.Add(n)

	// crear tenedores
	forks := make([]*tenedor, n)
	for i := 0; i < n; i++ {


		forks[i] = &tenedor{ID: i}
}
		

	}

	// lanzar filósofos
	for i := 0; i < n; i++ {
		izq := forks[i]
		der := forks[(i+1)%n]
		// TODO: lanzar goroutine para el filósofo i

	}

	wg.Wait()
	fmt.Println("Todos los filósofos han comido sin deadlock.")

	func aBinario(n int) string {

	if n == 0 {
		return "0"
	}
	
	binario := ""

	for n > 0 {
		residuo := n % 2
		binario = fmt.Sprintf("%d%s", residuo, binario)
		n = n / 2
	}

	return binario
}

