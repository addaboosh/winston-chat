package data

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Worker struct {
	Id       uuid.UUID                     `json:"id"`
	Name     string                        `json:"name"`
	Conns    map[uuid.UUID]*websocket.Conn `json:"conns"`
	Channels []string                      `json:"channels"`
	Groups   Group                         `json:"groups"`
}

func NewWorker() (wk *Worker, err error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return wk, err
	}
	return &Worker{
		Id:       id,
		Name:     "",
		Conns:    make(map[uuid.UUID]*websocket.Conn),
		Channels: []string{},
		Groups:   make(map[string][]uuid.UUID),
	}, nil
}

type Group map[string][]uuid.UUID
