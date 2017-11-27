package main

import "fmt"

func drawTab(tab []byte) {

	fmt.Println("				Fichas Blancas				")
	fmt.Println("    0    1    2    3    4    5    6    7   ")
	fmt.Println("  +----+----+----+----+----+----+----+----+")
	for i := 0; i < 8; i++ {
		fmt.Printf("%d ",i)
		for j := 0; j < 8; j++ {
			fmt.Printf("| %c  ", tab[i * 8 + j])
		}
		fmt.Println("|\n  +----+----+----+----+----+----+----+----+")
	}
	fmt.Println("				Fichas Negras				")
}

func scanJugada(tab []byte, p byte) {
	var aux byte
	var i, j int
	var k, l int
	validAll := false
	validMove := false
	for !validAll {
		if rune(p) == 'b' || rune(p) == 'B' {
			fmt.Printf("Seleccione ficha Blanca [0-1] [0-7]: ")
			fmt.Scanf("%d %d\n", &i, &j)
		} else  {
			//if rune(p) == 'n' || rune(p) == 'N'	
			fmt.Printf("Seleccione ficha Negra [6-7] [0-7]: ")
			fmt.Scanf("%d %d\n", &i, &j)
		}
		
		idx := i * 8 + j
		if i >= 0 && i < 8 && j >= 0 && j < 8 && tab[idx] != byte(' ') {
			validAll = true
			aux = tab[idx]
			for !validMove {
				fmt.Printf("Posición a  mover : ")
				fmt.Scanf("%d %d\n", &k, &l)
				
				idy := k * 8 + l
				if validateMove(tab, i, j, k ,l) == true {
					tab[idy] = aux
					tab[idx] = byte(' ')
					validMove = true
				} else {
					fmt.Println(" --- Jugada no permitida ---")
				}
			}
		} else {
			fmt.Println(" --- Ficha no permitida ---")
		}

	}
}

func validateMove(tab []byte, i int, j int, k int, l int) bool {
	var diffx, diffy int
	diffx = i - k
	diffy = j - l
	
	if rune(tab[i * 8 + j]) == 'p' || rune(tab[i * 8 + j]) == 'P' {	
		if (diffx == 1 || diffx == -1) && (diffy == 1 || diffy == -1) {
			if tab[k * 8 + l] != byte(' ') {
				return true
			}
		} else if (diffx == 1 || diffx == -1) && (diffy == 0) {
			if tab[k * 8 + l] == byte(' ') {
				return true
			}
		}
	}
	
	return false
}

func findWinner(tab []byte) byte {
	
	var contadorR int
	//var aux1,aux2 int 
	var contadorr int

	for i:= 0;i<64;i++ {

		if(tab[i]=='r'){
			contadorr = 1
	//		aux1 = i
		}
		if(tab[i]=='R'){
			contadorR = 1
			//aux2 = i
		}

	}
	if (contadorR == 0){
		return byte('N')
	}

	if (contadorr == 0){
		//return tab[aux2]
		return byte('B')
	}

	return 0

}

func chooseOpositeToken(color byte) byte {
	if color == byte('b') {
		return byte('n')
	} else if color == byte('B') {
		return byte('N')
	} else if color == byte('n') {
		return byte('b')
	} else if color == byte('N') {
		return byte('B')
	}
	return 0
}

func pickToken() byte {
	p := '-'
	for p != 'b' && p != 'n' && p != 'B' && p != 'N' {
		fmt.Println("Seleccione color [Blanco --> b,Negro -- > n]: ");
		fmt.Scanf(	"%c\n", &p)
		if p != 'b' &&  p != 'n' && p != 'B' &&  p != 'N' {
			fmt.Println(" --- Selección no permitida ---")
		}
	}
	
	return byte(p)
}

func getMsg(buff []byte) byte {
	return buff[0]
}
func getSessId(buff []byte) byte {
	return buff[1]
}
func getPlayerId(buff []byte) byte {
	return buff[2]
}
func getColor(buff []byte) byte {
	return buff[3]
}
func getTab(buff []byte) []byte {
	return buff[4:]
}
func setMsg(buff []byte, msg byte) {
	buff[0] = msg
}
func setSessId(buff []byte, sid byte) {
	buff[1] = sid
}
func setPlayerId(buff []byte, pid byte) {
	buff[2] = pid
}
func setColor(buff []byte, color byte) {
	buff[3] = color
}
func setTab(buff []byte, tab []byte) {
	for i, e := range tab {
		buff[i + 4] = e
	}
}

const (
	NEW = byte(0) // Mensajes del cliente
	UPDATE = byte(1)
	PLAY = byte(2)

	WAIT = byte(3)  // Mensajes del server
	TURN = byte(4)
)