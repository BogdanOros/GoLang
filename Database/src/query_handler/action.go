package query_handler

import (
	repo "./../database"
)

type Action interface {
	do (input [] string, repo *repo.DatabaseRepository) (string, error)
	validate (input [] string) bool
	execute (input [] string, repo *repo.DatabaseRepository) string
}

