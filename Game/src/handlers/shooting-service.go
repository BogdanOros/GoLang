package handlers

import (
	"../model"
	"../resources"
	"../services"
)

type ShootingResult struct {
	NeedToRepaint, ChangePlayer, RemoveHealth bool
}

func ProcessShootingProcess(clicked int, shipsHolder *model.ShipsHolder, board *model.Board, shootingBoard *model.Board) ShootingResult {
	switch clicked {
	case resources.Up:
		if shootingBoard.CurrentY != 0 {
			shootingBoard.CurrentY--
			return ShootingResult{true, false, false}
		}
		return ShootingResult{false, false, false}
	case resources.Down:
		if shootingBoard.CurrentY != resources.BoardSize - 1 {
			shootingBoard.CurrentY++
			return ShootingResult{true, false, false}
		}
		return ShootingResult{false, false, false}
	case resources.Left:
		if shootingBoard.CurrentX != 0 {
			shootingBoard.CurrentX--
			return ShootingResult{true, false, false}
		}
		return ShootingResult{false, false, false}
	case resources.Right:
		if shootingBoard.CurrentX != resources.BoardSize - 1 {
			shootingBoard.CurrentX++
			return ShootingResult{true, false, false}
		}
		return ShootingResult{false, false, false}
	case resources.Press:
		changePlayer, shipCell := services.CheckShootedCell(board, shootingBoard)
		if isShip(shipCell) {
			shipsHolder.ChangeShipsLife(shipCell)
		}
		return ShootingResult{true, !changePlayer, true}
	}
	return ShootingResult{false, false, false}
}

func isShip(shipCell int) bool {
	return shipCell > 0
}