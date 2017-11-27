package main 

import (
	"fmt"
	"net"
	"bufio"
)

func main() {
	conn, _ := net.Dial("tcp", "10.11.9.140:8080")
	defer conn.Close()	
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)
	buff := make([]byte, 1024)
	msgs := []string{"Hola, mundo", "Adios, mundo", "zzz"}
	for _, msg := range msgs {
		w.Write([]byte(msg))
		w.Flush()
		n, _ := r.Read(buff)
		fmt.Printf("Respuesta: %s\n", buff[:n])
	}

}