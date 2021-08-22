package saver

import (
	"fmt"
	"ova-conference-api/internal/domain"
	"ova-conference-api/internal/utils/flusher"
	"sync"
	"time"
)

type Saver interface {
	Save(conference domain.Conference)
	Init(duration time.Duration)
	Close()
}

type state int

const (
	empty state = iota
	initialized
	closed
)

type repository struct {
	flusher flusher.Flusher
	buffer  []domain.Conference
	*sync.Mutex
	timer       *time.Timer
	initializer *sync.Once
	finalizer   *sync.Once
	closeChanel chan bool
	state       state
}

func (rep *repository) Save(conference domain.Conference) {
	if rep.state != initialized {
		return
	}
	rep.Mutex.Lock()
	rep.buffer = append(rep.buffer, conference)
	rep.Mutex.Unlock()
}

func (rep *repository) Close() {
	if rep.state == initialized {
		rep.Lock()
		rep.state = closed
		rep.closeChanel <- true
		rep.Unlock()
	}
}

func (rep *repository) Init(duration time.Duration) {
	var initialF = func() {
		rep.timer = time.NewTimer(duration)
		rep.state = initialized
		go rep.awaitingFlush()
	}
	rep.initializer.Do(initialF)
}

func (rep *repository) flush() {
	rep.finalizer.Do(func() {
		errorEntities := rep.flusher.Flush(rep.buffer)
		if len(errorEntities) > 0 {
			fmt.Printf("error entities %v\n", errorEntities)
		}
		close(rep.closeChanel)
		rep.timer.Stop()
	})
}

func (rep repository) awaitingFlush() {
	for {
		select {
		case <-rep.closeChanel:
			rep.flush()
			return
		case <-rep.timer.C:
			rep.state = closed
			rep.flush()
			return
		}
	}
}

func NewSaver(capacity int, flusher flusher.Flusher) Saver {
	return &repository{
		flusher:     flusher,
		buffer:      make([]domain.Conference, capacity, capacity),
		Mutex:       &sync.Mutex{},
		initializer: &sync.Once{},
		finalizer:   &sync.Once{},
		closeChanel: make(chan bool),
		state:       empty,
	}
}
