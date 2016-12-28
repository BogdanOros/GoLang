package game

import (
	"../tools"
	"../model"
	"../handlers"
	"fmt"
	"../resources"
)

func GameLoop () {

	handler := handlers.KeyboardHandlerInit()
	go handler.WaitingForKey()

	session := gameProcessPreConfiguration()

	session.CurrentPlayer = &session.FirstPlayer
	shipPlacingProcess(&session, handler)

	session.CurrentPlayer = &session.SecondPlayer
	shipPlacingProcess(&session, handler)

	session.CurrentPlayer = &session.FirstPlayer

	tools.ClearScreen()
	tools.DrawBoard(session.CurrentPlayer.ShootingBoard)
	handleClicks(&session, handler)

	tools.ClearScreen()
	fmt.Println("Winner - ", session.CurrentPlayer.Name)
}

func handleClicks(session *GameSession, handler handlers.KeyboardHandler) {

	for {
		clicked := <- handler.GetKeyHandler()

		result :=
			handlers.ProcessShootingProcess(clicked, &session.CurrentPlayer.ShipsHolder, &session.CurrentPlayer.Board, &session.CurrentPlayer.ShootingBoard)

		if result.RemoveHealth {
			session.CurrentPlayer.Health -= resources.Damaged
		}

		if session.CurrentPlayer.Health == resources.Dead {
			break;
		}

		if result.ChangePlayer{
			changeCurrentPlayer(session)
		}

		if result.NeedToRepaint {
			tools.ClearScreen()
			tools.DrawBoard(session.CurrentPlayer.ShootingBoard)
			repaintInfo(session, "Press SPACE to shoot")
		}

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
	tools.DrawInterface(session.FirstPlayer, session.SecondPlayer, *session.CurrentPlayer, info)
}

func gameProcessPreConfiguration() GameSession{
	gameSession := GameSessionInit("First", "Second")
	return gameSession
}

func shipPlacingProcess(session *GameSession, handler handlers.KeyboardHandler) {
	for i:=0; i < len(session.CurrentPlayer.ShipsHolder.ShipsArray); i++ {
		session.CurrentPlayer.Board.ResetInitialPos()
		ship := session.CurrentPlayer.ShipsHolder.GetShip(i)
		session.CurrentPlayer.ShipsHolder.SetIndex(i)
		repaintBoard(&session.CurrentPlayer.Board, &ship)
		repaintInfo(session, "Ship placed")
		handleClicksWithShip(session, handler, &session.CurrentPlayer.Board, &ship)
	}
	session.CurrentPlayer.Readiness = true
}

func changeCurrentPlayer(session *GameSession) {
	if (*session.CurrentPlayer == session.FirstPlayer) {
		session.CurrentPlayer = &session.SecondPlayer
	} else {
		session.CurrentPlayer = &session.FirstPlayer
	}
	session.CurrentPlayer.ShootingBoard.ResetInitialPos()
}