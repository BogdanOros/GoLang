package model

import "../resources"

type Ship struct {
	Length int
	Direction int
	Alive bool
	Life int
	Index int
}

func ShipInit(len int) Ship {
	return Ship {
		Length: len,
		Life: len,
		Direction: resources.DIR_X,
		Alive: true,
	}
}

func (ship *Ship) SetIndex(index int) {
	ship.Index = index
}

func (ship *Ship) GetDamage() {
	if ship.Alive {
		ship.Life --
		if ship.Life == 0 {
			ship.Alive = false
		}
	}
}