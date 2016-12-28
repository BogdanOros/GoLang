package database

import (
	prop "./../resources"
	fileUtils "./../files"
	"fmt"
	"strings"
)

type DatabaseRepository struct {
	collections map[string] Collection
	activeCollection *Collection
}

func RepoInit () DatabaseRepository {
	return DatabaseRepository{collections: make(map[string] Collection, prop.DATABASES_COUNT)}
}

func (repo *DatabaseRepository) CreateCollection(collection string, fields [] string) Collection{
	coll := CollectionInit(collection, fields)
	repo.collections[collection] = coll
	return repo.collections[collection]
}

func (repo *DatabaseRepository) GetCollection(collection string) (Collection, error){
	coll, ok :=  repo.collections[collection];
	if !ok {
		fmt.Print(collection)
		res, err := fileUtils.ReadCollectionType(collection)
		if err == nil {
			return repo.CreateCollection(collection, strings.Fields(res)), nil
		}
		return coll, err
	}
	return coll, nil
}