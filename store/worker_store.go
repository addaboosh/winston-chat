package store

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Worker struct {
	Id          uuid.UUID    `json:"id"`
	Name        string       `json:"name"`
	Connections []Connection `json:"connections"`
}

type Connection struct {
	Id       uuid.UUID
	Conn     websocket.Conn
	Channels []string
}

type Group map[string][]uuid.UUID

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
}
