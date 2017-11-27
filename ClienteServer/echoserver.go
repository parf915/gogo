package main 

import (
	"fmt"
	"net"
	"bufio"
)

func servicio(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)
	buff := make([]byte, 1024)
	for {
		n, _ := r.Read(buff)
		fmt.Printf("Recibido: %s\n", buff[:n])
		if buff[0] != byte('z') {
			w.Write(buff)
			w.Flush()
		} else {
			fmt.Println("Bye bye!")
			break
		}
	}
}

func main() {
	lstnr, _ := net.Listen("tcp", "10.11.9.140:8080")
	defer lstnr.Close()
	for {
		conn, _ := lstnr.Accept()
		go servicio(conn)
	}
}