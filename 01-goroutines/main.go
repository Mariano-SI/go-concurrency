package main

import (
	"fmt"
	"time"
)

func main() {
	// exemplo 1:
	// Se executarmos as duas funções normalmente (sem "go"),
	// o programa ficará bloqueado na primeira chamada, pois ela nunca termina.
	// A segunda função nunca seria executada.
	// write("Hello, World!")
	// write("Programming in golang")

	// exemplo 2: (Esse funciona)
	// Ao adicionar "go" antes da chamada, criamos uma goroutine.
	// Isso faz com que a função seja executada de forma concorrente,
	// sem bloquear o fluxo principal do programa.
	// Nesse caso, a primeira função roda em paralelo ao restante do código.
	// go write("Hello, World!")
	// write("Programming in golang")

	// exemplo 3:
	// Se iniciarmos duas goroutines, ambas funções rodam de forma concorrente.
	// Porém, como a função main termina imediatamente, o programa encerra
	// antes que possamos ver a saída das goroutines.
	// (Para evitar isso, precisamos de alguma forma de sincronização
	// como WaitGroup, canais, ou até um simples Sleep na main).
	go write("Hello, World!")
	go write("Programming in golang")

	time.Sleep(5 * time.Second)
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
