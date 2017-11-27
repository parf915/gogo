package main 

import (
	"fmt"
	"net"
	"bufio"
)

func main() {
	ls, _ := net.Listen("tcp", "localhost:8080")
	conn, _ := ls.Accept()
	defer conn.Close()
	r := bufio.NewReader(conn)
	buf := make([]byte, 1024)
	n, _ := r.Read(buf)
	fmt.Printf("Bytes recibidos: %d\nMensaje:\n%s\n", n, buf[:n])

	msg := fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n\n%s\n\n\n",
		"HTTP/1.1 200 OK",
		"Server: nginx/1.13.4",
		"Date: Wed, 25 Oct 2017 16:59:05 GMT",
		"Content-Type: text/html",
		"Content-Length: 12",
		"Last-Modified: Tue, 12 Sep 2017 20:37:29 GMT",
		"Connection: keep-alive",
		"ETag: \"59b84589-dd\"",
		"Accept-Ranges: bytes",
		"Hola, Mundo!")
	w := bufio.NewWriter(conn)
	w.Write([]byte(msg))
	w.Flush()
}