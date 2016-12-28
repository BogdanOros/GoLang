package query_handler

import (
	repo "./../database"
	fileUtils "./../files"
	mappingUtils "./../entity"
	res "./../resources"
	"sync"
	"fmt"
)

type InsertAction struct {

	Mutex sync.Mutex
	Keyword [3]string
	MinLength int
}

func (action *InsertAction) do(input [] string, repos *repo.DatabaseRepository) (string, error) {
	action.Mutex.Lock()
	result, err := action.doUnsafe(input, repos)
	action.Mutex.Unlock()
	return result, err
}


func (action *InsertAction) doUnsafe(input [] string, repos *repo.DatabaseRepository) (string, error) {
	collection, crErr := repos.GetCollection(input[2])
	if crErr != nil {
		return crErr.Error(), nil
	}
	entity, err := mappingUtils.CreateMapEntity(collection, input[4:])
	if err != nil {
		return res.INSERT_NOT_VALID_ENTITY, nil
	}
	go fileUtils.SaveToFile(collection.CollectionName, entity)
	return res.INSERT_SUCCESS, nil
}

func (action InsertAction) validate(input [] string)  bool {
	fmt.Println(input)
	correctFormat :=  len(input) > action.MinLength &&
		input[0] == action.Keyword[0] &&
		input[1] == action.Keyword[1] &&
		input[3] == action.Keyword[2]
	return correctFormat
}

func (action InsertAction) execute (input [] string, repo *repo.DatabaseRepository) string {
	valid := action.validate(input)
	if valid {
		result, _ := action.do(input, repo)
		return result
	}
	return res.INSERT_NOT_VALID
}

