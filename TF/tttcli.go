package main

import (
	"fmt"
	"net"
	"bufio"
	"time"
)

func callServer(msg, sid, pid *byte, color *byte, tab *[]byte) {
	conn, _ := net.Dial("tcp", "localhost:8080")
	defer conn.Close()
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)
	buff := make([]byte, 68)
	setMsg(buff, *msg)
	setSessId(buff, *sid)
	setPlayerId(buff, *pid)
	setColor(buff, *color)
	setTab(buff, *tab)
	w.Write(buff)
	w.Flush()
	r.Read(buff)
	*msg = getMsg(buff)
	*sid = getSessId(buff)
	*pid = getPlayerId(buff)
	*color = getColor(buff)
	*tab = getTab(buff)
}

func main() {
	tab := make([]byte, 64)
	tab =  []byte{'T','C','A','Y','R','A','C','T',
				  'P','P','P','P','P','P','P','P',
				  ' ',' ',' ',' ',' ',' ',' ',' ',
				  ' ',' ',' ',' ',' ',' ',' ',' ',
				  ' ',' ',' ',' ',' ',' ',' ',' ',
				  ' ',' ',' ',' ',' ',' ',' ',' ',
				  'p','p','p','p','p','p','p','p',
				  't','c','a','y','r','a','c','t'}
				  
	var msg, sid, pid, color, winner byte
	var p byte
	msg = NEW
	firstTurn := true
	for winner == 0 {
		callServer(&msg, &sid, &pid, &color, &tab)
		winner = findWinner(tab)
		switch msg {
		case WAIT:
			time.Sleep(time.Second)
			msg = UPDATE
		case TURN:
			if firstTurn {
				firstTurn = false
				p = chooseOpositeToken(color)
				if p == 0 {
					p = pickToken()
					color = p
				}
			}
			fmt.Println("Tablero actual: ")
			drawTab(tab)
			if winner == 0 {
				scanJugada(tab, p)
				drawTab(tab)
				msg = PLAY
			}
		}
	}
	fmt.Printf("Ganaron las %c!\n", winner)
}