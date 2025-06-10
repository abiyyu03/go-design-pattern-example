package generator

import (
	"fmt"
	"sync"
	"time"
)

type (
	Event struct {
		data int
	}

	Observer interface {
		NotifyCallback(Event)
	}

	Subject interface {
		AddListener(Observer)
		RemoveListener(Observer)
		Notify(Event)
	}

	EventObserver struct {
		id   int
		time time.Time
	}

	EventSubject struct {
		observers sync.Map
	}
)

func (e *EventObserver) NotifyCallback(event Event) {
	fmt.Printf("Observer: %d Recieved: %d after %v\n", e.id, event.data, time.Since(e.time))
}

func (s *EventSubject) AddListener(obs Observer) {
	s.observers.Store(obs, struct{}{})
}

func (s *EventSubject) RemoveListener(obs Observer) {
	s.observers.Delete(obs)
}

func (s *EventSubject) Notify(event Event) {
	s.observers.Range(func(key interface{}, value interface{}) bool {
		if key == nil || value == nil {
			return false
		}

		key.(Observer).NotifyCallback(event)
		return true
	})
}

func Fib(n int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}

	}()
	return out
}
