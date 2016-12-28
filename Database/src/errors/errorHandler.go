package errors

import (
	"os"
	"fmt"
	"net"
)

func HandleErrorDefault (err error, errorMessage string) {
	if err != nil {
		fmt.Println(errorMessage, err.Error())
		os.Exit(1)
	}
}

func HandleErrorClient (err error, errorMessage string, conn net.Conn) {
	if err != nil {
		fmt.Println(errorMessage, err.Error())
		conn.Write([]byte(errorMessage + "\n"))
	}
}