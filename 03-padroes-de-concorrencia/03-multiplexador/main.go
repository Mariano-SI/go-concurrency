package main

import (
	"fmt"
	"time"
)

func main() {
	// Chamamos a função multiplexar, que recebe dois canais de entrada
	// e retorna um canal unificado que entrega os valores de ambos.
	channel := multiplexar(write("Olá, Mundo"), write("Trabalhando com go"))

	// Consumimos 10 valores do canal multiplexado
	for i := 0; i < 10; i++ {
		message := <-channel
		fmt.Println(message)
	}
}

// Função multiplexador: combina vários canais em um único canal de saída
func multiplexar(channel1, channel2 <-chan string) <-chan string {
	// Criamos o canal de saída unificado
	channel := make(chan string)

	// Goroutine que ficará constantemente lendo ambos os canais de entrada
	go func() {
		for {
			select {
			// Se channel1 tiver valor, envia para o canal unificado
			case message := <-channel1:
				channel <- message
			// Se channel2 tiver valor, envia para o canal unificado
			case message := <-channel2:
				channel <- message
			}
		}
	}()

	// Retornamos o canal multiplexado
	return channel
}

// Função generator: retorna um canal que gera valores continuamente
func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			// Envia o valor para o canal
			channel <- fmt.Sprintf("Valor recebido %s", text)
			// Aguarda meio segundo antes de gerar o próximo valor
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return channel
}
