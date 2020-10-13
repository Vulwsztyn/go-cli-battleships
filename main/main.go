package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const mapSize = 10

type cell struct {
	shot         bool
	containsShip bool
}

type Board = [][]cell

func initBoard() (board Board) {
	board = make(Board, mapSize, mapSize)

	for i := 0; i < mapSize; i++ {
		board[i] = make([]cell, mapSize, mapSize)
		//for j := 0; j < mapSize; j++ {
		//	board[i][j] = cell{shot: false, containsShip: false}
		//}
	}
	return
}

func addShip(r1 *rand.Rand, board Board, len int) Board {
	for unplaced := true; unplaced; {
		horizontal := r1.Intn(2) > 0
		rowColNum := r1.Intn(10)
		offsetInRowCol := r1.Intn(11 - len)
		positions := make([][2]int, len)
		for i := 0; i < len; i++ {
			if horizontal {
				positions[i] = [2]int{rowColNum, offsetInRowCol + i}
			} else {
				positions[i] = [2]int{offsetInRowCol + i, rowColNum}
			}
		}
		arePositionsUnoccupied := true
		for _, pos := range positions {
			if board[pos[0]][pos[1]].containsShip {
				arePositionsUnoccupied = true
				break
			}
		}
		if arePositionsUnoccupied {
			for _, pos := range positions {
				board[pos[0]][pos[1]].containsShip = true
			}
			unplaced = false
		}
	}
	return board
}

func printBoard(board Board, debug bool) {
	fmt.Print(" ")
	for i := 0; i < mapSize; i++ {
		fmt.Print(" ", string(rune(65+i)))
	}
	fmt.Println()
	for i := 0; i < mapSize; i++ {
		fmt.Print(i)
		for j := 0; j < mapSize; j++ {
			toBePrinted := "| "
			if board[i][j].shot || debug {
				if board[i][j].containsShip {
					toBePrinted = "|X"
				} else {
					toBePrinted = "|."
				}
			}
			fmt.Print(toBePrinted)
		}
		fmt.Println("|")
	}
}

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	aiBoard := initBoard()
	toBeShot := 0
	for i := 4; i > 1; i-- {
		aiBoard = addShip(r1, aiBoard, i)
		toBeShot += i
	}

	for ; toBeShot > 0; {
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		inputText := input.Text()
		if inputText == "show" {
			printBoard(aiBoard, true)
		} else {
			col := int(inputText[0]) - int(rune('A'))
			row := int(inputText[1]) - int(rune('0'))
			if aiBoard[row][col].shot {
				fmt.Println("Already shot")
			} else {
				aiBoard[row][col].shot = true
				if aiBoard[row][col].containsShip {
					toBeShot -= 1
				}
			}
		}
		printBoard(aiBoard, false)
		fmt.Println(toBeShot)
	}
}
