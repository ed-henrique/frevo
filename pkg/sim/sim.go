package sim

import (
	"fmt"
	"time"

	"github.com/ed-henrique/frevo/internal/queues"
	"github.com/ed-henrique/frevo/pkg/event"
)

const (
	warningScheduledAfterStopTime = "tried to schedule an event after the simulation stop time."
)

type Simulation struct {
	q              *queues.EventQueue
	stopTime       time.Duration
	currentTime    time.Duration
	enableWarnings bool
}

type option func(*Simulation)

func WithStopTime(t time.Duration) option {
	return func(s *Simulation) {
		s.stopTime = t
	}
}

func WithWarnings() option {
	return func(s *Simulation) {
		s.enableWarnings = true
	}
}

func New(options ...option) *Simulation {
	s := &Simulation{q: queues.NewEventQueue()}

	for _, o := range options {
		o(s)
	}

	return s
}

func (s *Simulation) CurrentTime() time.Duration {
	return s.currentTime
}

func (s *Simulation) Run() {
	var stop func() bool

	if s.stopTime == 0 {
		stop = func() bool { return true }
	} else {
		stop = func() bool { return s.currentTime <= s.stopTime }
	}

	for stop() && s.q.Len() > 0 {
		e, timestamp := s.q.Pop()

		if s.stopTime != 0 && timestamp > s.stopTime {
			break
		}

		s.currentTime = timestamp
		e.Do()
	}
}

func (s *Simulation) Schedule(t time.Duration, e event.Event) {
	if s.enableWarnings {
		if s.stopTime != 0 && t > s.stopTime {
			fmt.Println("WARNING: ", warningScheduledAfterStopTime)
		}
	}

	s.q.Push(e, t)
}
