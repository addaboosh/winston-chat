package main

import (
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

func (i *Instance) AddWorker(w *Worker) {
	i.workers[w.id] = *w
}
func (i *Instance) RemoveWorker(id uuid.UUID){
	// ToDo - Need to check if any active connections before deleting
	delete(i.workers, id)
}
