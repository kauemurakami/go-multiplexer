package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplexy(write("Hello world"), write("Go lang"))

	for k := 0; k < 10; k++ {
		fmt.Println(<-channel)
	}
}

// ao chamar esssa função mais de uma vez você terá mais de um canal
// a ideia é juntar esses canais em um só canal para centralizar
// a comunicação em um lugar só na main
// canal que apenas recebe valores
func multiplexy(channelEnter1, channelEnter2 <-chan string) <-chan string {
	channelOutput := make(chan string)

	//goroutine anonymous
	go func() {
		//for infinito
		for {
			// vai verificar qual dos dois canais vai ter
			// uma mensagem disponivel pra ser lida
			// se for no canal1 ele vai jogar pro canal de saida
			// se for no canal2 também vai jogar pro canal de saida
			select {
			case message := <-channelEnter1:
				channelOutput <- message
			case message := <-channelEnter2:
				channelOutput <- message
			}
		}
	}()
	return channelOutput
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Value received: %s", text)
			//valor de delay aleatorio entre 1 e 2 sec
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
	return channel
}
