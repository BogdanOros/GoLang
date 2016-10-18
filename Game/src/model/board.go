package model

type Board struct {
	board [10][10] int
	CurrentX, CurrentY int
}

func BoardInit() Board {
	return Board{}
}

func (board Board) GetCell(x, y int) int {
	return board.board[x][y]
}

