package main

import (
	"fmt"
)

type TestEvent struct {
	ID string
}

func (e TestEvent) EventName() string {
	return "TestEvent"
}

type HandlerOne struct{}

func (h HandlerOne) Handle(event Event) {
	fmt.Println("handler1")
	fmt.Println(event.EventName())
}

type HandlerTwo struct{}

func (h HandlerTwo) Handle(event Event) {
	fmt.Println("handler2")
	fmt.Println(event.EventName())
}

func main() {
	dispatcher := &Dispatcher{handlers: make(map[string][]EventHandler)}
	testEvent := TestEvent{ID: "xxxx"}
	handlerOne := HandlerOne{}
	handlerTwo := HandlerTwo{}

	dispatcher.Subscribe(testEvent, handlerOne)
	dispatcher.Subscribe(testEvent, handlerTwo)
	dispatcher.Publish(testEvent)
}
