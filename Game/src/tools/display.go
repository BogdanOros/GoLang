package tools

import (
	"../model"
	"fmt"
	"../resources"
)

func DrawInterface(firstPlayer, secondPlayer model.Player, gameInfo string) {
	GoToXY(11, 0);
	fmt.Println("CurrentGameInformation: ", gameInfo);
	fmt.Println("First player - ", firstPlayer.Name, " : ", firstPlayer.Readiness)
	fmt.Println("Second player - ", secondPlayer.Name, " : ", secondPlayer.Readiness)
}

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
	case resources.DIR_X:
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if j >= board.CurrentX && j < board.CurrentX + ship.Length && board.CurrentY == i{
					fmt.Print(" * ")
					continue
				}
				fmt.Print(getSymbolByCell(board.GetCell(i, j)))
			}
			fmt.Println()
		}
	case resources.DIR_Y:
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if i >= board.CurrentY && i < board.CurrentY + ship.Length && board.CurrentX == j {
					fmt.Print(" * ")
					continue
				}
				fmt.Print(getSymbolByCell(board.GetCell(i, j)))
			}
			fmt.Println()
		}
	}
}

func getSymbolByCell(cell int) string {
	switch cell {
	case 0:
		return " . "
	case 1:
		return " # "
	default:
		return "   "
	}
}





