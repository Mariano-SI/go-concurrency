package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// WaitGroup serve para garantir que todas as goroutines terminem
	// antes do programa se encerrar
	var waitGroup sync.WaitGroup

	// Vamos monitorar 2 goroutines
	waitGroup.Add(2)

	// Primeira goroutine
	go func() {
		write("Hello, World!")
		waitGroup.Done() // sinaliza que terminou (-1)
	}()

	// Segunda goroutine
	go func() {
		write("Programming in Go")
		waitGroup.Done() // sinaliza que terminou (-1)
	}()

	// Espera até que todas as goroutines chamem Done()
	waitGroup.Wait()
}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
