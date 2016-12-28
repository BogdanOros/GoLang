package entity

import (
	coll "./../database"
	"errors"
)

func CreateMapEntity(collection coll.Collection, input[] string) (map[string] string, error) {
	if len(collection.Fields) != len(input) {
		return nil, errors.New("Incorrect insertion value")
	}
	entity := make(map[string] string)
	for i, element := range collection.Fields {
		entity[element] = input[i]
	}
	return entity, nil
}
