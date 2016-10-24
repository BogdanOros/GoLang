package model

type Board struct {
	Board [10][10] int
	CurrentX, CurrentY int
}

func BoardInit() Board {
	return Board{}
}

func (board Board) GetCell(x, y int) int {
	return board.Board[x][y]
}

func (board *Board) ResetInitialPos() {
	board.CurrentX = 0
	board.CurrentY = 0
}

