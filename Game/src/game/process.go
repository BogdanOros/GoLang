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

	session := gameProcessPreConfiguration()

	shipPlacingProcess(&session, handler)

	board1 := model.BoardInit()
	tools.ClearScreen()
	tools.DrawBoard(board1)
	handleClicks(handler, &board1)

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

func handleClicksWithShip(session *GameSession, handler handlers.KeyboardHandler, board *model.Board, ship *model.Ship) {
	for {
		clicked := <- handler.GetKeyHandler()
		needToRepaint, shipPlaced := handlers.ProcessShipMovingKbHit(clicked, board, ship)
		if needToRepaint && !shipPlaced{
			repaintBoard(board, ship)
			repaintInfo(session, "Ship placing process")
		} else if needToRepaint && shipPlaced {
			repaintBoard(board, ship)
			repaintInfo(session, "Ship placing process finished")
			break
		}
	}
}

func repaintBoard(board *model.Board, ship *model.Ship) {
	tools.ClearScreen()
	tools.DrawBoardWithShipMoving(*board, *ship)
}

func repaintInfo(session *GameSession, info string) {
	tools.DrawInterface(session.FirstPlayer, session.SecondPlayer, info)
}

func gameProcessPreConfiguration() GameSession{
	gameSession := GameSessionInit("First", "Second")
	gameSession.SetCurrentPlayer(&gameSession.FirstPlayer)
	return gameSession
}

func shipPlacingProcess(session *GameSession, handler handlers.KeyboardHandler) {
	currentUser := session.GetCurrentPlayer()
	for i:=0; i < len(currentUser.ShipsHolder.ShipsArray); i++ {
		currentUser.Board.ResetInitialPos()
		ship := currentUser.ShipsHolder.GetShip(i)
		repaintBoard(&currentUser.Board, &ship)
		repaintInfo(session, "Ship placed")
		handleClicksWithShip(session, handler, &currentUser.Board, &ship)
	}
	currentUser.Readiness = true
}

