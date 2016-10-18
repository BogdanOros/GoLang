package services

import (
	"../model"
	"../resources"
	"fmt"
)

func MoveUp(board *model.Board, ship *model.Ship) bool {
	if board.CurrentY != getBorder(resources.Up, ship) {
		board.CurrentY--
		return true
	}
	return false
}

func MoveDown(board *model.Board, ship *model.Ship) bool {
	if board.CurrentY != getBorder(resources.Down, ship) {
		board.CurrentY++
		return true
	}
	return false
}

func MoveLeft(board *model.Board, ship *model.Ship) bool {
	if board.CurrentX != getBorder(resources.Left, ship) {
		board.CurrentX--
		return true
	}
	return false

}
func MoveRight(board *model.Board, ship *model.Ship) bool {
	if board.CurrentX != getBorder(resources.Right, ship) {
		board.CurrentX++
		return true
	}
	return false
}
func ChangeDirection(board *model.Board, ship *model.Ship) bool{

	border := getBorder(resources.Press, ship)

	switch ship.Direction {
	case resources.DIR_X:
		if board.CurrentY <= border {
			fmt.Print(ship.Direction)
			ship.Direction = resources.DIR_Y
			return true
		}
		return false
	case resources.DIR_Y:
		if board.CurrentX <= border {
			fmt.Print(ship.Direction)
			ship.Direction = resources.DIR_X
			return true
		}
		return false
	}
	return false
}

func getBorder(clicked int, ship *model.Ship) int{
	switch clicked {
	case resources.Up:
		switch ship.Direction {
		case resources.DIR_X:
			return 0
		case resources.DIR_Y:
			return 0
		}
	case resources.Down:
		switch ship.Direction {
		case resources.DIR_X:
			return resources.BoardSize - 1
		case resources.DIR_Y:
			return resources.BoardSize - 1 - ship.Length
		}
	case resources.Left:
		switch ship.Direction {
		case resources.DIR_X:
		case resources.DIR_Y:
			return  0
		}
	case resources.Right:
		switch ship.Direction {
		case resources.DIR_X:
			return resources.BoardSize - 1 - ship.Length
		case resources.DIR_Y:
			return resources.BoardSize - 1
		}
	case resources.Press:
		switch ship.Direction {
		case resources.DIR_X:
			return resources.BoardSize - 1 - ship.Length
		case resources.DIR_Y:
			return resources.BoardSize - 1 - ship.Length
		}
	}
	return resources.ERROR
}

