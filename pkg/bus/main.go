package bus

import (
	"context"

	"github.com/LSN0WM4N/filessh/pkg/env"
)

type EventType string

const (
	EventKey    EventType = "key"
	EventResize EventType = "resize"
	EventFocus  EventType = "focus"
	EventSelect EventType = "select"
	EventQuit   EventType = "quit"
	EventAction EventType = "action"
)

type Event struct {
	Type    EventType
	Payload any
}

type KeyInfo struct {
	Key rune
	Seq string
}

type ResizeInfo struct {
	Width  int
	Height int
}

type Handler func(Event)

type EventBus struct {
	ch       chan Event
	handlers map[EventType][]Handler
}

func NewEventBus() *EventBus {
	return &EventBus{
		ch:       make(chan Event, env.GetIntEnv("EVENT_QUEUE_SIZE", 64)),
		handlers: make(map[EventType][]Handler),
	}
}

// Register an Handler to a key
//
//	bus.Subscribe(Event, func(e Event) {
//	    data := e.Payload
//	    // do stuff...
//	})
func (b *EventBus) Subscribe(t EventType, fn Handler) {
	b.handlers[t] = append(b.handlers[t], fn)
}

// Publish an event to the bus.
//
// IMPORTANT: By default, there are only 64 events in the bus, if
// you call this function when the event queue is full, the main thread
// will be frozen until the new event is added to the queue, to
// increase the queue size recompile with flag EVENT_QUEUE_SIZE=N
// where N is the new maximum
func (b *EventBus) Publish(e Event) {
	b.ch <- e
}

// Dispatcher. Must run on it's own goroutine:
//
//	go bus.Run(ctx)
func (b *EventBus) Run(ctx context.Context) {
	for {
		select {

		case e := <-b.ch:
			for _, fn := range b.handlers[e.Type] {
				fn(e)
			}

			if e.Type == EventQuit {
				return
			}

		case <-ctx.Done():
			return
		}
	}
}
