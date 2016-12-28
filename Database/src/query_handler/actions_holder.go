package query_handler

type ActionHolder struct {
	actions map[string] interface {}
}

func ActionHolderInit () ActionHolder {
	actions := map[string] func (a Action) {
		"create" : CreateAction{"create", 3}.execute,
	}
	return ActionHolder{
	 	actions: actions,
	}
}
