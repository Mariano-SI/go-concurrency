package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//wait group serve para garantir que todas as routines irao ter terminado antes do porgrama se encerrar

	var waitGroup sync.WaitGroup

	waitGroup.Add(2) // quantidade de routines que ele vai monitorar

	go func(){
		write("Hello, World!")
		waitGroup.Done() //-1
	}()
	go func(){
		write("Programming in golang")
		waitGroup.Done()//-1
	}()
	
	waitGroup.Wait() // permite que o programa escerre quando a quantidade for 0
}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
