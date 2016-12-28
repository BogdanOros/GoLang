package services

import ("../resources"; "../model";
)

func CheckShootedCell(board *model.Board, shootingBoard *model.Board) (bool, int) {
	cell := board.GetCell(shootingBoard.CurrentY, shootingBoard.CurrentX)
	stayAtCurrentPlayer := true
	switch cell {
	case resources.NoShipCell:
		shootingBoard.Board[shootingBoard.CurrentY][shootingBoard.CurrentX].BoardCell = resources.ShootedNoShipCell
		stayAtCurrentPlayer = false
	case resources.ShipCell:
		shootingBoard.Board[shootingBoard.CurrentY][shootingBoard.CurrentX].BoardCell = resources.ShootedShipCell
		stayAtCurrentPlayer = true
	default:
		stayAtCurrentPlayer =  true
	}
	return stayAtCurrentPlayer, board.Board[shootingBoard.CurrentY][shootingBoard.CurrentX].ShipCell
}
