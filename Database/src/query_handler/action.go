package query_handler

type Action interface {
	do (input [] string)
	validate (input [] string) bool
	execute (input [] string)
}

