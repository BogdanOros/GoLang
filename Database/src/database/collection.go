package database

type Collection struct {
	collectionName string
	fields [] string
}

func CollectionInit(collName string, fields [] string) *Collection {
	return &Collection{collName, fields}
}