package main

import "fmt"

const MAX = 9
const K = 4

func compress(inc, pipe chan rune) {
	n := 0
	previous := <- inc
	for c := range inc {
		if c == previous && n < MAX-1 {
			n++
		} else {
			if n > 0 {
				pipe <- rune('0' + n + 1)
				n = 0
			}
			pipe <- previous
			previous = c
		}
	}
	if n > 0 {
		pipe <- rune('0' + n + 1)
	}
	pipe <- previous
	close(pipe)
}
func output(pipe, outc chan rune) {
	m := 0
	for c := range pipe {
		outc <- c
		m++
		if m >= K {
			outc <- '\n'
			m = 0
		}
	}
	close(outc)
}
func main() {
	inc := make(chan rune)
	pipe := make(chan rune)
	outc := make(chan rune)
	go compress(inc, pipe)
	go output(pipe, outc)
	go func() {
		cadena := "abccccdddeeeffghh"
		for _, letra := range cadena {
			inc <- letra
		}
		close(inc)
	}()
	for c := range outc {
		fmt.Printf("%c", c);
	}
	fmt.Println("")
}