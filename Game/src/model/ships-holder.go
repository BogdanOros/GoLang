package model


type ShipsHolder struct {
	ShipsArray [8] Ship
}

func ShipsHolderInit() ShipsHolder {
	var shipsArray [8] Ship = [8] Ship {
		ShipInit(4),
		ShipInit(3),
		ShipInit(3),
		ShipInit(2),
		ShipInit(2),
		ShipInit(1),
		ShipInit(1),
	}
	return ShipsHolder{ShipsArray: shipsArray}
}

func (holder ShipsHolder) GetShip(index int) Ship {
	return holder.ShipsArray[index]
}

func (holder *ShipsHolder) SetIndex(index int ) {
	holder.ShipsArray[index].SetIndex(index + 1)
}

func (holder *ShipsHolder) ChangeShipsLife(index int) *Ship {
	ship := &holder.ShipsArray[index]
	ship.GetDamage()
	return ship
}