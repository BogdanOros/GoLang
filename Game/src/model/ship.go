package model

import "../resources"

type Ship struct {
	Length int
	Direction int
	Alive bool
}

func ShipInit() Ship {
	return Ship {
		Length: 3,
		Direction: resources.DIR_X,
		Alive: true,
	}
}