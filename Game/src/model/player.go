package model

type Player struct {
	Name string
	Board, ShootingBoard Board
	ShipsHolder ShipsHolder
	Health int
	Readiness bool
}

func PlayerInit(name string) Player {
	shipHolder := ShipsHolderInit()
	health := 0
	for _, ship := range shipHolder.ShipsArray {
		health += ship.Length
	}

	return Player {
		Name : name,
		Board: BoardInit(),
		ShootingBoard: BoardInit(),
		ShipsHolder: shipHolder,
		Health: health,
	}
}