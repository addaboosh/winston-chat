package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/addaboosh/winston-chat/data"
)

type Worker struct {
	l *log.Logger
}

func NewWorker(l *log.Logger) *Worker{
	return &Worker{l}
}

func (w *Worker) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello from worker handler")
	i := data.GetInstance()
	data := i.GetWorkers()
	enc := json.NewEncoder(rw)
	enc.Encode(data)
	w.l.Println(data)
}
