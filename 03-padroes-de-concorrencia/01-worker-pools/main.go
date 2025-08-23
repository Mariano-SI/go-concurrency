package main

import "fmt"

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-1) + fibonacci(posicao-2)
}

func main() {
	tasks := make(chan int, 45)
	results := make(chan int, 45)

	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i <= 45; i++ {
		tasks <- i
	}

	close(tasks)

	for i := 0; i <= 45; i++ {
		result := <-results
		fmt.Println(result)
	}
}

func worker(tasks <-chan int, results chan<- int) {
	for task := range tasks {
		results <- fibonacci(task)
	}
}
