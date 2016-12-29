package query_handler

import (
	repo "./../database"
	//fileUtils "./../files"
	res "./../resources"
	entityUtils "./../entity"
	"sync"
)

type SelectAction struct {
	Mutex sync.Mutex
	Keyword [3]string
	MinLength int
}

func (action *SelectAction) do(input [] string, repos *repo.DatabaseRepository) (string, error) {
	action.Mutex.Lock()
	result, err := action.doUnsafe(input, repos)
	action.Mutex.Unlock()
	return result, err
}

func (action *SelectAction) doUnsafe(input [] string, repos *repo.DatabaseRepository) (string, error) {
	unformattedCollection, err := repos.ReadCollection(input[2])
	if (err != nil) {
		return err.Error(), nil
	}
	curCol, err := repos.GetCollection(input[2])
	if (err != nil) {
		return err.Error(), nil
	}
	formattedCollection := entityUtils.CreateMapEntityList(curCol, unformattedCollection)
	result := entityUtils.FormatEntityToString(curCol, formattedCollection)
	return result, nil
}

func (action SelectAction) validate(input [] string)  bool {
	correctFormat :=  len(input) > action.MinLength &&
		input[0] == action.Keyword[0] &&
		input[1] == action.Keyword[1]
	return correctFormat
}

func (action SelectAction) execute (input [] string, repo *repo.DatabaseRepository) string {
	valid := action.validate(input)
	if valid {
		result, _ := action.do(input, repo)
		return result
	}
	return res.CREATE_NOT_VALID
}


