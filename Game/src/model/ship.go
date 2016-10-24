package model

import "../resources"

type Ship struct {
	Length int
	Direction int
	Alive bool
}

func ShipInit(len int) Ship {
	return Ship {
		Length: len,
		Direction: resources.DIR_X,
		Alive: true,
	}
}