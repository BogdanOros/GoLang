package services

import (
	"../model"
	"../resources"
)

func PlaceShip(board *model.Board, ship *model.Ship) {
	switch ship.Direction {
	case resources.DIR_X:
		for i := board.CurrentX; i < board.CurrentX + ship.Length; i++ {
			board.Board[board.CurrentY][i].BoardCell = 1
			board.Board[board.CurrentY][i].ShipCell = ship.Index
		}
	case resources.DIR_Y:
		for j := board.CurrentY; j < board.CurrentY + ship.Length; j++ {
			board.Board[j][board.CurrentX].BoardCell = 1
			board.Board[j][board.CurrentX].ShipCell = ship.Index
		}
	}
}

func CheckPlacingAvailability(board *model.Board, ship *model.Ship) bool{
	isAbleToPlace := true
	switch ship.Direction {
	case resources.DIR_X:
		for i := board.CurrentX; i < board.CurrentX + ship.Length && isAbleToPlace; i++ {
			if board.Board[board.CurrentY][i].BoardCell == 1 {
				isAbleToPlace = false
			}
		}
	case resources.DIR_Y:
		for j := board.CurrentY; j < board.CurrentY + ship.Length && isAbleToPlace; j++ {
			if board.Board[j][board.CurrentX].BoardCell == 1 {
				isAbleToPlace = false
			}
		}
	}
	return isAbleToPlace
}
