package main

import "fmt"

func main() {
	// Criando um canal com buffer de capacidade 2
	// - Um canal **sem buffer** bloqueia a goroutine enviadora até que alguém leia.
	// - Um canal **com buffer** permite enviar até N mensagens antes de bloquear.
	// Nesse exemplo, como o buffer é 2 e só enviamos 1 mensagem, não há bloqueio.
	//
	// Quando usar canais com buffer:
	// - Quando você quer permitir que a goroutine enviadora continue trabalhando
	//   mesmo que a receptora ainda não esteja pronta para receber todas as mensagens.
	// - Útil para desacoplar produção e consumo de dados, evitando que o envio fique bloqueado imediatamente.
	channel := make(chan string, 2)

	// Envia uma mensagem para o canal.
	channel <- "Hello, World!" 

	// Recebe a mensagem do canal.
	message := <-channel

	fmt.Println(message)
}
