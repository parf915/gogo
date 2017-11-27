package main 

import (
	"fmt"
	"net"
	"bufio"
)

func main() {
	lstnr, _ := net.Listen("tcp",
		"localhost:8080")
	conn, _ := lstnr.Accept()
	defer conn.Close()
	buff := make([]byte, 1024)
	r := bufio.NewReader(conn)
	n, _ := r.Read(buff)
	fmt.Printf("bytes leidos: %d\nMsg: %s\n", n, buff[:n])
}