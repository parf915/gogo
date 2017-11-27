package main

import "fmt"

func main() {
	ch := make(chan string)
	ch <- "5"
	// ch <- 6 // produce deadlock
	x := <- ch
	// x = <- ch // produce deadlock
	fmt.Printf("Valor leído: %s\n", x)
}