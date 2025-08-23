package main

import "fmt"

// Função que calcula o número de Fibonacci de forma recursiva
func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-1) + fibonacci(posicao-2)
}

func main() {
	// Criamos dois canais:
	// - tasks: canal de tarefas que contém os números para calcular o Fibonacci
	// - results: canal de resultados, onde cada worker envia o resultado do cálculo
	tasks := make(chan int, 45)
	results := make(chan int, 45)

	// Inicializamos várias goroutines "workers"
	// Cada worker vai pegar uma tarefa do canal "tasks", processar e enviar para "results"
	// Ter múltiplos workers permite processar várias tarefas **concorrentemente**
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	// Enviamos as tarefas para o canal "tasks"
	for i := 0; i <= 45; i++ {
		tasks <- i
	}

	// Fechamos o canal de tarefas, indicando que não haverá mais tarefas
	close(tasks)

	// Lemos os resultados do canal "results"
	// Cada cálculo feito por qualquer worker aparecerá aqui
	for i := 0; i <= 45; i++ {
		result := <-results
		fmt.Println(result)
	}
}

// Função worker:
// - Recebe tarefas do canal tasks
// - Calcula fibonacci da tarefa
// - Envia o resultado para o canal results
func worker(tasks <-chan int, results chan<- int) {
	for task := range tasks {
		results <- fibonacci(task)
	}
}
