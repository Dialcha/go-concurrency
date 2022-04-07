package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//Creamos 20 hilos de la función
	for i := 0; i < 20; i++ {
		go thread(i)
	}

	//Este tiempo es importante debido a que si el hilo
	//principal termina antes, los hilos no se
	//ejecutarán
	time.Sleep(101 * time.Second)
}

func thread(index int) {

	//Para simular una carga de trabajo
	//dormimos el programa x cantidad de segundo
	//donde x puede ir de x a 100
	var seconds int
	seconds = rand.Intn(100)
	time.Sleep(time.Duration(seconds) * time.Second)

	fmt.Println("Este es el hilo número", index, "(", seconds, "s )")
}
