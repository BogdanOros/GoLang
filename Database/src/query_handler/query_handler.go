package query_handler

import (
	"fmt"
	repo "./../database"
	"net"
)

func BuildQuery(input [] string, repository *repo.DatabaseRepository, conn net.Conn) {
	actions := ActionHolderInit()
	result := actions.execute(input, repository, conn)
	fmt.Println("ACTION RESULT: ", result)
}
