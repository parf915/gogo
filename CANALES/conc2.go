package main

import "fmt"

func P(id int, fin chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Printf("P%d: %d\n", id, i)
	}
	fin <- true
}

func main() {
	n := 5		
	fin := make(chan bool)
	for i := 0; i < n; i++ {
		go P(i, fin)
	}
	for i := 0; i < n; i++ {
		<- fin
	}
}