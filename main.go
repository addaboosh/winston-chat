package main

import (
	"github.com/addaboosh/winston-chat/handlers"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", handlers.NewHello())

	http.ListenAndServe(":8123", mux)
}

type Instance struct {
	l *log.Logger
	w map[uuid.UUID]Worker
}

type Worker struct {
	id       uuid.UUID
	name     string
	conns    map[uuid.UUID]*websocket.Conn
	channels []string
	groups   Group
}

func NewWorker() *Worker {
	return {
		id uuid.UUID. 

	}
}

type Group map[string][]uuid.UUID
