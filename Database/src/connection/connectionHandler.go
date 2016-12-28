package connection

import (
	"net"
	"bufio"
	"strings"
	prop "./../resources"
	errorHandler "./../errors"
	handler "./../query_handler"
	repo "./../database"
)

func StartConnection() {
	listener := createListener()
	repo := repo.RepoInit()
	defer listener.Close()
	for {
		connection, err := listener.Accept()
		errorHandler.HandleErrorDefault(err, "Error during new client acception.")
		go handleConnection(connection, &repo)
	}
}

func createListener() net.Listener {
	listenerAddress := prop.CONNECTION_HOST + ":" + prop.CONNECTION_PORT
	listener, err := net.Listen(prop.CONNECTION_TYPE, listenerAddress)
	errorHandler.HandleErrorDefault(err, "Error during listening port.")
	return listener
}

func handleConnection(conn net.Conn, repository *repo.DatabaseRepository) {
	for {
		buffer := bufio.NewReader(conn)
		for {
			query, err := buffer.ReadString('\n')
			errorHandler.HandleErrorClient(err, "Error during message reading.", conn)
			go handleQuery(conn, query, repository)
		}

	}
}

func handleQuery(conn net.Conn, query string, repository *repo.DatabaseRepository) {
	commands := strings.Fields(query)
	handler.BuildQuery(commands, repository, conn)
}