package handlers

import (
	"../model"
	"../services"
	"../resources"
)

/*
	@param bool : check need we repaint the board
	@param bool : check is the ship placed
*/
func ProcessShipMovingKbHit(clicked int, board *model.Board, ship *model.Ship) (bool, bool) {
	switch clicked {
	case resources.Up:
		return services.MoveUp(board, ship), false
	case resources.Down:
		return services.MoveDown(board, ship), false
	case resources.Left:
		return services.MoveLeft(board, ship), false
	case resources.Right:
		return services.MoveRight(board, ship), false
	case resources.Press:
		return services.ChangeDirection(board, ship), false
	case resources.Place:
		return ProcessShipPlacing(board, ship), true
	}
	return false, false
}