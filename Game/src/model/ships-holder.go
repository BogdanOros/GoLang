package model


type ShipsHolder struct {
	ShipsArray [8] Ship
}

func ShipsHolderInit() ShipsHolder {
	var shipsArray [8] Ship
	shipsArray[0] = ShipInit(4)
	shipsArray[1] = ShipInit(3)
	shipsArray[2] = ShipInit(3)
	shipsArray[3] = ShipInit(2)
	shipsArray[4] = ShipInit(2)
	shipsArray[5] = ShipInit(1)
	shipsArray[6] = ShipInit(1)
	return ShipsHolder{ShipsArray: shipsArray}
}

func (holder ShipsHolder) GetShip(index int) Ship {
	return holder.ShipsArray[index]
}
