package main

import "fmt"

func producer(id int, ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("Producer%d -> producto%d", id, i)
	}
}

func consumer(id int, ch chan string, fin chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Consumidor%d <- (%s)\n", id, <- ch)
	}
	fin <- true
}

func main() {
	ch := make(chan string)
	fin := make(chan bool)
	n := 5
	for i := 0; i < n; i++ {
		go producer(i, ch)
		go consumer(i, ch, fin)
	}
	for i := 0; i < n; i++ {
		<- fin
	}
}