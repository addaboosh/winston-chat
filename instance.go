package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type Instance struct {
	l       *log.Logger
	workers map[uuid.UUID]Worker
}

func NewInstance(l *log.Logger) *Instance {
	return &Instance{l, map[uuid.UUID]Worker{}}
}

func (i *Instance) GetWorker(id uuid.UUID) (*Worker, error) {

	w, ok := i.workers[id]
	if ok {
		return &w, nil
	}
	return nil, errors.New(fmt.Sprintf("Could not find Worker id %v", id))
}

func (i *Instance) CreateWorker() (id uuid.UUID, err error) {
	w, err := NewWorker()
	if err != nil {
		return uuid.UUID{}, err
	}
	i.workers[w.id] = *w
	return w.id, nil
}

func (i *Instance) RemoveWorker(id uuid.UUID) error {
	// ToDo - Need to check if any active connections before deleting
	delete(i.workers, id)
	_, ok := i.workers[id]
	if ok { 
		return nil
	}
	return errors.New(fmt.Sprintf("Worker id: %v not removed", id))
}
