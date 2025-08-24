package main

import (
	"fmt"
	"time"
)

func main() {
	// Chamamos a função write, que retorna um canal.
	// Esse canal será usado para receber os valores gerados pela goroutine interna.
	channel := write("Hello, World")

	// Consumimos os valores gerados pelo canal.
	// Cada leitura do canal recebe o próximo valor produzido pelo generator.
	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

// Função generator: retorna um canal que gera valores sob demanda
func write(text string) <-chan string {
	// Criamos o canal de saída
	channel := make(chan string)

	// Goroutine interna que gera valores infinitamente
	go func() {
		for {
			// Envia o valor para o canal
			channel <- fmt.Sprintf("Valor recebido %s", text)
			// Aguarda meio segundo antes de gerar o próximo valor
			time.Sleep(time.Millisecond * 500)
		}
	}()

	// Retornamos o canal para o consumidor
	return channel
}
