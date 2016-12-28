package model

import "../resources"

type Cell struct {
	ShipCell, BoardCell int
}

type Board struct {
	Board [resources.BoardSize][resources.BoardSize] Cell
	CurrentX, CurrentY int
}

func BoardInit() Board {
	return Board{}
}

func (board Board) GetCell(x, y int) int {
	return board.Board[x][y].BoardCell
}

func (board *Board) ResetInitialPos() {
	board.CurrentX = 0
	board.CurrentY = 0
}

