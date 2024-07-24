package store

import (
	"github.com/google/uuid"
)

type Worker struct {
	Id          uuid.UUID    `json:"id"`
	Name        string       `json:"name"`
	Connections map[uuid.UUID]IConnection `json:"connections"`
}

type CreateWorkerParams struct {
	Name string
}

type SetWorkerNameParams struct {
	Name string
}

type Interface interface {
	GetAll() ([]Worker, error)
	GetByID(id uuid.UUID) (Worker, error)
	Create(createWorkerParams CreateWorkerParams) (Worker, error)
	SetName(id uuid.UUID, setWorkerNameParams SetWorkerNameParams) error
	Delete(id uuid.UUID) error
	//CreateConnection(createConnectionParams CreateConnectionParams, err error) (int, error)
	//RemoveConnection(id uuid.UUID, err error) error
}
