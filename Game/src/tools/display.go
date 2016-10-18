package tools

import (
	"../model"
	"fmt"
)

func DrawBoard (board model.Board) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if board.CurrentX == j && board.CurrentY == i {
				fmt.Print(" _ ")
			} else {
				fmt.Print(" . ")
			}

		}
		fmt.Println()
	}
}

func DrawBoardWithShipMoving (board model.Board, ship model.Ship) {
	switch ship.Direction {
	case 0:
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if j >= board.CurrentX && j <= board.CurrentX + ship.Length && board.CurrentY == i{
					fmt.Print(" * ")
				} else {
					fmt.Print(" . ")
				}
			}
			fmt.Println()
		}
	case 1:
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if i >= board.CurrentY && i <= board.CurrentY + ship.Length && board.CurrentX == j {
					fmt.Print(" * ")
				} else {
					fmt.Print(" . ")
				}
			}
			fmt.Println()
		}
	}

}





