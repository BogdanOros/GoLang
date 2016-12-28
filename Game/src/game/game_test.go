package game

import "testing"
import (
	"../handlers"
	"../model"
	"../services"
	"fmt"
)

func createBoard(value int) *model.Board {
	board := model.BoardInit()
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			board.Board[i][j].BoardCell = value
		}
	}
	return &board
}

func Test001(t *testing.T) {
	board := createBoard(1)
	ship := model.ShipInit(5)
	if handlers.ProcessShipPlacing(board, &ship) {
		t.Error("Cant place a ship")
	}
}

func Test002(t *testing.T) {
	board := createBoard(0)
	ship := model.ShipInit(5)
	if !handlers.ProcessShipPlacing(board, &ship) {
		t.Error("Cant place a ship")
	}
}

func Test003(t *testing.T) {
	board := createBoard(0)
	ship := model.ShipInit(5)
	if services.MoveUp(board, &ship) {
		t.Error("Ship can`t move up")
	} else {
		fmt.Print("Ship can move up")
	}
}

func Test004(t *testing.T) {
	board := createBoard(0)
	ship := model.ShipInit(5)
	if services.MoveDown(board, &ship) {
		fmt.Println("Ship can move down")
	} else {
		t.Error("Ship can`t move down")
	}
}

func Test005(t *testing.T) {
	board := createBoard(0)
	ship := model.ShipInit(5)
	if services.MoveLeft(board, &ship) {
		t.Error("Ship can`t move left")
	} else {
		fmt.Println("Ship can move left")
	}
}


func Test006(t *testing.T) {
	board := createBoard(0)
	ship := model.ShipInit(5)
	if services.MoveRight(board, &ship) {
		fmt.Println("Ship can move right")
	} else {
		t.Error("Ship can`t move right")
	}
}