package main

import (
	"fmt"
	"time"
	"math/rand"
)

func escribe(id int, ch chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Nanosecond*time.Duration(rand.Intn(50)))
		ch <- i*id
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go escribe(1, ch1)
	go escribe(2, ch2)
	var val int
	for i := 0; i < 10; i++ {
		select {
			case val = <- ch1:
				fmt.Printf("Valor recibido de canal 1: %d\n", val)
			case val = <- ch2:
				fmt.Printf("Valor recibido de canal 2: %d\n", val)
		}
	}
}