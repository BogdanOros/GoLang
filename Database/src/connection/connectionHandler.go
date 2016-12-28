package connection

import (
	"net"
	"bufio"
	"strings"
	prop "./../resources"
	errorHandler "./../errors"
	handler "./../query_handler"
)

func StartConnection() {
	listener := createListener()
	defer listener.Close()
	for {
		connection, err := listener.Accept()
		errorHandler.HandleErrorDefault(err, "Error during new client acception.")
		go handleConnection(connection)
	}
}

func createListener() net.Listener {
	listenerAddress := prop.CONNECTION_HOST + ":" + prop.CONNECTION_PORT
	listener, err := net.Listen(prop.CONNECTION_TYPE, listenerAddress)
	errorHandler.HandleErrorDefault(err, "Error during listening port.")
	return listener
}

func handleConnection(conn net.Conn) {
	for {
		buffer := bufio.NewReader(conn)
		for {
			query, err := buffer.ReadString(prop.BUFFER_DELIMETER)
			errorHandler.HandleErrorClient(err, "Error during message reading.", conn)
			go handleQuery(conn, query)
		}

	}
}

func handleQuery(conn net.Conn, query string) {
	commands := strings.Fields(query)
	handler.BuildQuery(commands)
}