package data

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var instance = NewInstance()

type Instance struct {
	workers map[uuid.UUID]Worker
}

func GetInstance() *Instance {
	return instance
} 

func NewInstance() *Instance {
	return &Instance{map[uuid.UUID]Worker{}}
}

func (i *Instance) GetWorkers() *map[uuid.UUID]Worker{
	return &i.workers
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
	i.workers[w.Id] = *w
	return w.Id, nil
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
