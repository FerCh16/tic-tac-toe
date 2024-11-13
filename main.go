package main

import (
	"fmt"
	"errors"
	"strconv"
)
/*

0: Free
1: X
2: O

*/

type Table [3][3]int

var ErrNotAvailable = errors.New("Box not Available")

const (
	Free = iota
	Player1
	Player2
)

var Symbols = map[int]string {
	Free: "·",
	Player1: "X",
	Player2: "O",
}

func (t *Table) Display() {
	for _, row := range t {
		for _, v := range row {
			fmt.Printf(" %v ",Symbols[v])
		}
		fmt.Println()
	}
}

func (t *Table) Set(x, y, value int) error {

	if t[y-1][x-1] == Free {
		t[y-1][x-1] = value	
		return nil	
	}

	return ErrNotAvailable

}

func main() {

	var table Table

	var input string

	var x,y int64

	var IsPlayer1 = true


	fmt.Print("\033[2J\033[H")

	for {
		fmt.Println("TIC TAC TOE\n")
		table.Display()
		
		fmt.Print("Ingrese Coordenadas X,Y -> ")

		fmt.Scanf("%s",&input)

		x,_ = strconv.ParseInt(string(input[0]), 10, 0)
		y,_ = strconv.ParseInt(string(input[2]), 10, 0)

		var err error

		if IsPlayer1 {
			err = table.Set(int(x),int(y), Player1)
		}else {
			err = table.Set(int(x),int(y), Player2)
		}
	
		if err != nil {
			fmt.Println("Por favor ingrese una coordenada valida!\n")
		}

		IsPlayer1 = !IsPlayer1

		fmt.Print("\033[2J\033[H")


	}

}

