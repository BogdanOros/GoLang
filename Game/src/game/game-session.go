package game

import (
	"../model"
	"../resources"
)

type GameSession struct {
	FirstPlayer, SecondPlayer model.Player
	CurrentPlayer *model.Player
	GameState int
}

func GameSessionInit(player1, player2 string) GameSession {
	return GameSession {
		FirstPlayer: model.PlayerInit(player1),
		SecondPlayer: model.PlayerInit(player2),
		CurrentPlayer: new(model.Player),
		GameState: resources.ShipPlacingState,
	}
}

func (session *GameSession) GetCurrentPlayer() model.Player {
	return *session.CurrentPlayer
}

func (session *GameSession) SetCurrentPlayer(player *model.Player) {
	*session.CurrentPlayer = *player
}