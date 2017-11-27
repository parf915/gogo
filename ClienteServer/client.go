package main 

import (
	"fmt"
	"net"
	"bufio"
)

func main() {
	conn, _ := net.Dial("tcp",
		"localhost:8080")
	defer conn.Close()	
	w := bufio.NewWriter(conn)
	msg := "Hola, server!\n"
	w.Write([]byte(msg))
	w.Flush()
	fmt.Printf("bytes enviados: %d\nMsg: %s\n", len(msg), msg)
}