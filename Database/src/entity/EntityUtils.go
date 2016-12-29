package entity

import (
	coll "./../database"
	"errors"
	//"strings"
	"fmt"
	"encoding/json"
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

func createEntityFromSlice(input string) (map[string] string) {
	entity := make(map[string] string)
	json.Unmarshal([]byte(input), &entity)
	return entity
}

func CreateMapEntityList(collection coll.Collection, input[] string) []map[string] string {
	result := []map[string] string {}
	fmt.Println(input)
	for _, elem := range (input) {
		entity := createEntityFromSlice(elem)
		result = append(result, entity)
	}
	return result
}

func FormatEntityToString(collection coll.Collection, res []map[string]string) string {
	result := ""
	for _, elem := range(res) {
		for _, field := range(collection.Fields) {
			result += field + " : " + elem[field] + " ; "
		}
		result += " \n"
	}
	return result
}

func DeleteEntitiesFromArray(source []map[string]string, field, query string)  []map[string]string {
	result := [] map [string] string {}
	for _, elem := range(source) {
		if elem[field] != query {
			result = append(result, elem)
		}
	}
	return result
}
