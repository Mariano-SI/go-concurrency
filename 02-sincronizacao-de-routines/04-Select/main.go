package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	// Goroutine que envia mensagens a cada 0,5s para c1
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			c1 <- "Channel 1"
		}
	}()

	// Goroutine que envia mensagens a cada 2s para c2
	go func() {
		for {
			time.Sleep(time.Second * 2)
			c2 <- "Channel 2"
		}
	}()

	// Problema do loop sequencial:
	// Se fizermos assim:
	/*
	for {
		messageC1 := <-c1
		fmt.Println(messageC1)

		messageC2 := <-c2
		fmt.Println(messageC2)
	}
	*/
	// A goroutine principal fica **bloqueada** esperando c2.
	// Enquanto c2 ainda não envia (demora 2s), c1 poderia ter enviado várias mensagens,
	// mas não conseguimos processá-las imediatamente. Isso prejudica a performance
	// e o consumo contínuo de mensagens de c1.

	// Jeito correto usando select:
	// O select permite reagir a **qualquer canal que esteja pronto**.
	// Assim, conseguimos processar mensagens de c1 e c2 assim que chegam,
	// sem bloquear o processamento de um canal esperando o outro.
	for {
		select {
		case messageC1 := <-c1:
			fmt.Println(messageC1)
		case messageC2 := <-c2:
			fmt.Println(messageC2)
		}
	}
}
