package main 

import (
	"fmt"
	"net"
	"bufio"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:8080")
	defer conn.Close()
	w := bufio.NewWriter(conn)

	msg := fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n\n\n",
		"GET / HTTP/1.1",
		"Host: localhost:8080",
		"Connection: keep-alive",
		"Cache-Control: max-age=0",
		"User-Agent: Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
		"Upgrade-Insecure-Requests: 1",
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
		"Accept-Encoding: gzip, deflate, br",
		"Accept-Language: es-ES,es;q=0.8")
	w.Write([]byte(msg))
	w.Flush()
	r := bufio.NewReader(conn)
	buf := make([]byte, 1024)
	n, _ := r.Read(buf)
	fmt.Printf("Bytes recibidos: %d\nMensaje:\n%s\n", n, buf[:n])
}