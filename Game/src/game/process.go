package game

import (
	"../tools"
	"../model"
	"../handlers"
	"../resources"
)

func GameLoop () {

	handler := handlers.KeyboardHandlerInit()
	go handler.WaitingForKey()

	tools.ClearScreen()

	board := model.BoardInit()
	ship := model.ShipInit()
	tools.DrawBoardWithShipMoving(board, ship)

	handleClicksWithShip(handler, &board, &ship)
	//handleClicks(handler, &board)
}

func handleClicks(handler handlers.KeyboardHandler, board *model.Board) {
	for {
		clicked := <- handler.GetKeyHandler()
		switch clicked {
		case resources.Up:
			if board.CurrentY != 0 {
				board.CurrentY--
			}
		case resources.Down:
			if board.CurrentY != 9 {
				board.CurrentY++
			}
		case resources.Left:
			if board.CurrentX != 0 {
				board.CurrentX--
			}
		case resources.Right:
			if board.CurrentX != 9 {
				board.CurrentX++
			}
		}
		tools.ClearScreen()
		tools.DrawBoard(*board)
	}
}

func handleClicksWithShip(handler handlers.KeyboardHandler, board *model.Board, ship *model.Ship) {
	for {
		clicked := <- handler.GetKeyHandler()
		needToRepaint := handlers.ProcessShipMovingKbHit(clicked, board, ship)
		if needToRepaint {
			tools.ClearScreen()
			tools.DrawBoardWithShipMoving(*board, *ship)
		}
	}
}

