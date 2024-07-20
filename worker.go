package main

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)


type Worker struct {
	id       uuid.UUID
	name     string
	conns    map[uuid.UUID]*websocket.Conn
	channels []string
	groups   Group
}

func NewWorker() (wk *Worker, err error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return wk, err
	}
	return &Worker{
		id: id,
		name: "",
		conns: make(map[uuid.UUID]*websocket.Conn),
		channels: []string{},
		groups: make(map[string][]uuid.UUID),

	}, nil
}

type Group map[string][]uuid.UUID
