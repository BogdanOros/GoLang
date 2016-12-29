package query_handler

import (
	repos "./../database"
	"net"
	res "./../resources"
)

type ActionHolder struct {
	actions map[string] interface {}
}

func ActionHolderInit () ActionHolder {
	actions := map[string] interface{} {
		"create" : CreateAction{Keyword: [2]string{"create", "as"}, MinLength: 3}.execute,
		"insert" : InsertAction{Keyword: [3]string{"insert", "into", "values"}, MinLength: 4}.execute,
		"select" : SelectAction{Keyword: [3]string{"select", "from" , "where"}, MinLength: 2}.execute,
	}
	return ActionHolder{
	 	actions: actions,
	}
}

func (holder ActionHolder) execute (input [] string, repo *repos.DatabaseRepository, conn net.Conn) bool {
	command := holder.actions[input[0]]
	if command == nil {
		conn.Write([]byte(res.INCORRECT_COMMAND + res.BUFFER_DELIMETER))
		return false
	}
	result := command.(func([]string, *repos.DatabaseRepository) string)(input, repo)
	conn.Write([]byte(result + "\n"))
	return true
}
