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
	closeChanel chan struct{}
	state       state
}

func (rep *repository) Save(conference domain.Conference) {
	rep.Mutex.Lock()
	defer rep.Mutex.Unlock()

	if rep.state != initialized {
		return
	}
	rep.buffer = append(rep.buffer, conference)
}

func (rep *repository) Close() {
	rep.Mutex.Lock()
	defer rep.Mutex.Unlock()

	if rep.state != initialized {
		return
	}
	rep.state = closed
	rep.timer.Stop()
	close(rep.closeChanel)
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
	if len(rep.buffer) > 0 {
		errorEntities := rep.flusher.Flush(rep.buffer)
		if len(errorEntities) > 0 {
			fmt.Printf("error entities %v\n", errorEntities)
		}
	}
}

func (rep repository) awaitingFlush() {
	for {
		select {
		case <-rep.closeChanel:
			rep.flush()
			return
		case <-rep.timer.C:
			rep.state = closed
			close(rep.closeChanel)
		}
	}
}

func NewSaver(capacity int, flusher flusher.Flusher) Saver {
	return &repository{
		flusher:     flusher,
		buffer:      make([]domain.Conference, capacity, capacity),
		Mutex:       &sync.Mutex{},
		initializer: &sync.Once{},
		closeChanel: make(chan struct{}),
		state:       empty,
	}
}
