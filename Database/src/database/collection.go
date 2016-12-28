package database

type Collection struct {
	CollectionName string
	Fields [] string
}

func CollectionInit(collName string, fields [] string) Collection {
	return Collection{collName, fields}
}