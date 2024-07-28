package main

type Event interface {
	EventName() string
}

type EventHandler interface {
	Handle(event Event)
}

type Dispatcher struct {
	handlers map[string][]EventHandler
}

func (d *Dispatcher) Subscribe(event Event, handler EventHandler) {
	d.handlers[event.EventName()] = append(d.handlers[event.EventName()], handler)
}

func (d *Dispatcher) Publish(event Event) {
	handlers := d.handlers[event.EventName()]
	for _, handler := range handlers {
		handler.Handle(event)
	}
}
