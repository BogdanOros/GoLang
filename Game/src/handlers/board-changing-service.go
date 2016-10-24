package handlers

import (
	"../model"
	"../services"
)

func ProcessShipPlacing(board *model.Board, ship *model.Ship) bool {
	isAbleToPlace := services.CheckPlacingAvailability(board, ship)
	if isAbleToPlace {
		services.PlaceShip(board, ship)
		return true
	}
	return false
}