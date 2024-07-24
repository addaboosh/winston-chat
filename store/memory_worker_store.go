package store

import (
	"sync"

	"github.com/google/uuid"
)

type MemoryWorkerStore struct {
	workers map[uuid.UUID]Worker
	mu      sync.RWMutex
}

func NewMemoryWorkerStore() *MemoryWorkerStore {
	return &MemoryWorkerStore{
		workers: map[uuid.UUID]Worker{},
	}
}

func (s *MemoryWorkerStore) GetAll() ([]Worker, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var workers []Worker
	for _, w := range s.workers {
		workers = append(workers, w)
	}
	return workers, nil
}

func (s *MemoryWorkerStore) GetByID(id uuid.UUID) (Worker, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	w, ok := s.workers[id]
	if !ok {
		return Worker{}, &ResourceNotFoundError{}
	}
	return w, nil
}

func (s *MemoryWorkerStore) Create(createWorkerParams CreateWorkerParams) (Worker, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return Worker{}, &UUIDCreateError{}
	}
	worker := Worker{
		Id: newUUID,
		Name: createWorkerParams.Name,
		Connections: map[uuid.UUID]IConnection{},
	}
	s.workers[worker.Id] = worker
	return s.workers[worker.Id], nil 
}

func(s *MemoryWorkerStore) SetName(id uuid.UUID, setWorkerNameParams SetWorkerNameParams) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	w, ok := s.workers[id]
	if !ok {
		return &ResourceNotFoundError{}
	}

	w.Name = setWorkerNameParams.Name

	s.workers[id] = w
	return nil
}

func (s *MemoryWorkerStore) Delete(id uuid.UUID) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.workers, id)
	return nil
}
