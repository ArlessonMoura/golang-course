package main

import (
	"fmt"
	"sync"
)

// Variável de nível de pacote: Decide o numero de rodadas "ping-pong"
const totalRodadas = 10 

func ping(c chan string, wg *sync.WaitGroup) {
	defer wg.Done() 
	
	for i := 0; i < totalRodadas; i++ { 
		msg := <-c
		fmt.Println(msg)
		c <- "pong"
	}
}

func pong(c chan string, wg *sync.WaitGroup) {
	defer wg.Done() 
	
	for i := 0; i < totalRodadas; i++ { 
		msg := <-c
		fmt.Println(msg)
		c <- "ping"
	}
}

func main() {
	var wg sync.WaitGroup
	// Cria um canal para comunicação entre as goroutines
	c := make(chan string) 

	wg.Add(2) 

	go ping(c, &wg)
	go pong(c, &wg)

	c <- "ping" 

	wg.Wait() 
	fmt.Println("Fim do jogo!")
}