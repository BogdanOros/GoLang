package query_handler

import (
	repo "./../database"
	fileUtils "./../files"
	res "./../resources"
	"sync"
)

type CreateAction struct {
	Mutex sync.Mutex
	Keyword [2]string
	MinLength int
}

func (action *CreateAction) do(input [] string, repos *repo.DatabaseRepository) (string, error) {
	action.Mutex.Lock()
	result, err := action.doUnsafe(input, repos)
	action.Mutex.Unlock()
	return result, err
}

func (action *CreateAction) doUnsafe(input [] string, repos *repo.DatabaseRepository) (string, error) {
	collection := repos.CreateCollection(input[1], input[3:])
	go fileUtils.CreateFile(collection.CollectionName, collection.Fields)
	return res.CREATE_SUCCESS, nil
}

func (action CreateAction) validate(input [] string)  bool {
	correctFormat :=  len(input) > action.MinLength &&
		input[0] == action.Keyword[0] &&
		input[2] == action.Keyword[1]
	return correctFormat
}

func (action CreateAction) execute (input [] string, repo *repo.DatabaseRepository) string {
	valid := action.validate(input)
	if valid {
		result, _ := action.do(input, repo)
		return result
	}
	return res.CREATE_NOT_VALID
}


