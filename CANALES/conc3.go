package main

import "fmt"

func P(id int, ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("P%d: %d", id, i)
	}
}

func main() {
	n := 5		
	ch := make(chan string)
	for i := 0; i < n; i++ {
		go P(i, ch)
	}
	for i := 0; i < n*10; i++ {
		fmt.Println(<- ch)
	}
}