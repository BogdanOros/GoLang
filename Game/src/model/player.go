package model

type Player struct {
	Name string
	Board Board
	ShipsHolder ShipsHolder
	Readiness bool
}

func PlayerInit(name string) Player {
	return Player {
		Name : name,
		Board: BoardInit(),
		ShipsHolder: ShipsHolderInit(),
	}
}