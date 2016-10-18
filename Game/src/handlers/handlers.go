package handlers

import (
	"../tools"
)

type KeyboardHandler struct {
	handler chan int
}

func KeyboardHandlerInit () KeyboardHandler {
	return KeyboardHandler {handler: make(chan int)}
}

func (handler KeyboardHandler) WaitingForKey () {
	for {
		handler.handler <- tools.ReadKey()
	}
}

func (handler KeyboardHandler) GetKeyHandler() chan int {
	return handler.handler
}


