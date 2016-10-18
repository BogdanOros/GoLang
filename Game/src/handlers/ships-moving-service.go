package handlers

import (
	"../model"
	"../services"
	"../resources"
)

func ProcessShipMovingKbHit(clicked int, board *model.Board, ship *model.Ship) bool {
	switch clicked {
	case resources.Up:
		return services.MoveUp(board, ship)
	case resources.Down:
		return services.MoveDown(board, ship)
	case resources.Left:
		return services.MoveLeft(board, ship)
	case resources.Right:
		return services.MoveRight(board, ship)
	case resources.Press:
		return services.ChangeDirection(board, ship)
	}
	return false
}