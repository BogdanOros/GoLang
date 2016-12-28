package database

import prop "./../resources"

type DatabaseRepository struct {
	collections [] Collection
	head int
}

func RepoInit () DatabaseRepository {
	return DatabaseRepository{make([]Collection, prop.DATABASES_COUNT), 0}
}

func (repo *DatabaseRepository) CreateCollection(collection string, fields [] string) {
	coll := CollectionInit(collection, fields)
	repo.collections[repo.head] = coll
	repo.head += 1
}
