package model

type Ship struct {
	Length int
	Direction int
	Alive bool
}

func ShipInit() Ship {
	return Ship {
		Length: 3,
		Direction: 0,
		Alive: true,
	}
}