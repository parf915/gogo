package main 

import "fmt"

type Tripleta struct {
	i, j int
	val float32
}

func zero(n int, west chan float32) {
	for i := 0; i < n; i++ {
		west <- 0
	}
	close(west)
}
func sink(north chan float32) {
	for {
		<- north
	}
}
func result(i int, east chan float32, res chan Tripleta) {
	j := 0
	for val := range east {
		res <- Tripleta{i, j, val}
		j++
	}
}
func source(row []float32, south chan float32) {
	for _, val := range row {
		south <- val
	}
	close(south)
}
func multiplier(n int, firstElement float32,
			    north, east, south, west chan float32) {
	var secondElement, sum float32
	for i := 0; i < n; i++ {
		select {
		case secondElement = <- north:
			sum = <- east
		case sum = <- east:
			secondElement = <- north
		}
		select {
		case west <- sum + firstElement*secondElement:
			south <- secondElement
		case south <- secondElement:
			west <- sum + firstElement*secondElement
		}
	}
	close(west)
	close(south)
}

func matmul(a, b [][]float32) [][]float32 {
	rowsA := len(a)
	colsA := len(a[0])
	rowsB := len(b)
	colsB := len(b[0])
	if colsA != rowsB {
		return nil
	}
	res := make(chan Tripleta)
	ns := make([][]chan float32, rowsA+1)
	for i := 0; i < rowsA+1; i++ {
		ns[i] = make([]chan float32, colsA)
		for j := 0; j < colsA; j++ {
			ns[i][j] = make(chan float32)
		}
	}
	ew := make([][]chan float32, rowsA)
	for i := 0; i < rowsA; i++ {
		ew[i] = make([]chan float32, colsA+1)
		for j := 0; j < colsA+1; j++ {
			ew[i][j] = make(chan float32)
		}
	}
	for i := 0; i < rowsA; i++ {
		go zero(colsB, ew[i][colsA])
		go result(i, ew[i][0], res)
	}
	for i := 0; i < rowsB; i++ {
		go sink(ns[rowsA][i])
		go source(b[i], ns[0][i])
	}
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsA; j++ {
			go multiplier(rowsA, a[i][j],
				          ns[i][j], ew[i][j+1],
				          ns[i+1][j], ew[i][j])
		}
	}
	c := make([][]float32, rowsA)
	for i := 0; i < rowsA; i++ {
		c[i] = make([]float32, colsB)
	}
	for i := 0; i < rowsA*colsB; i++ {
		r := <- res
		c[r.i][r.j] = r.val
	}
	return c
}
func main() {
	a := [][]float32{{1, 2, 3},
					 {4, 5, 6}}
	b := [][]float32{{1, 0},
	                 {0, 1},
	                 {1, 0}}
	fmt.Printf("%5.1f", matmul(a, b))
}