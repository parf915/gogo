package main

import (
	"fmt"
)

func multiplier(firstElement float32,
			    north, east, south, west chan float32) {
	var secondElement, sum float32
	for {
		select {
			case secondElement = <- north:
				sum = <- east
			case sum = <- east
				secondElement = <- north
		}
		south <- secondElement
		west <- sum + firstElement*secondElement
	}
}
func zero(n int, west chan float32) {
	for i := 0; i < n; i++ {
		west <- 0
	}
}
func source(row []float32, south chan float32) {
	for _, e := range row {
		south <- e
	}	
}
func sink(n int, north chan float32) {
	for i := 0; i < n; i++ {
		<- north
	}
}
func result(n int, east chan float32, fin chan bool) {
	for i := 0; i < n; i++ {
		fmt.Printf("%6.1f", <- east)
	}
	fmt.Println("")
	fin <- true
}

func main() {
	n := 3
	a := [][]float32{{1, 2, 3},
					 {4, 5, 6},
					 {7, 8, 9}}
	b := [][]float32{{1, 0, 2},
					 {0, 1, 2},
					 {1, 0, 0}}
	fin := make(chan bool)
	ns := make([][]chan float32, n + 1)
	ew := make([][]chan float32, n)
	for i := 0; i < n + 1; i++ {
		ns[i] = make([]chan float32, n)
		for j := 0; j < n; j++ {
			ns[i][j] = make(chan float32)
		}
	}
	for i := 0; i < n; i++ {
		ew[i] = make([]chan float32, n + 1)
		for j := 0; j < n + 1; j++ {
			ew[i][j] = make(chan float32)
		}
	}
	for i := 0; i < n; i++ {
		go zero(n, ew[i][n])
		go source(b[i], ns[0][i])
		go sink(n, ns[n][i])
		go result(n, ew[i][0], fin)
		for j := 0; j < n; j++ {
			go multiplier(a[i][j],
						  ns[i][j], ew[i][j+1],
						  ns[i+1][j], ew[i][j])
		}
	}
	for i := 0; i < n; i++ {
		<- fin
	}
}
