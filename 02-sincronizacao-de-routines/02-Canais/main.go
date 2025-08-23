package main

import (
	"fmt"
	"time"
)

func main() {
	// O canal é a forma idiomática de sincronizar goroutines em Go.
	// Ele permite a comunicação segura entre diferentes execuções concorrentes,
	// enviando e recebendo valores.
	canal := make(chan string)

	// Executa a função `write` em uma goroutine separada, enviando mensagens pelo canal.
	go write("Hello, World!", canal)

	// Loop que consome as mensagens enviadas pelo canal usando a sintaxe "for range".
	// Cada iteração recebe automaticamente o valor enviado pelo canal.
	// O loop termina automaticamente quando o canal é fechado, evitando a necessidade
	// de verificar manualmente se o canal ainda está aberto.
	for message := range canal {
		fmt.Println(message)
	}

	fmt.Println("Fim do programa")
}

// Função concorrente que escreve repetidamente no canal.
// Repare que a função recebe o canal como parâmetro, permitindo comunicação direta.
func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		// Envia o valor para o canal. Caso não haja nenhum receptor pronto,
		// a goroutine ficará bloqueada até que alguém leia.
		channel <- text
		time.Sleep(time.Second)
	}

	// É fundamental fechar o canal ao final.
	// Isso sinaliza para o receptor que não haverá mais mensagens,
	// evitando deadlocks e loops infinitos.
	close(channel)
}
