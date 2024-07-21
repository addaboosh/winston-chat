package main

import (
	"log"
	"net/http"
	"os"

	"github.com/addaboosh/winston-chat/data"
	"github.com/addaboosh/winston-chat/handlers"
)

func main() {
		instance := data.GetInstance()	
	instance.CreateWorker()
	instance.CreateWorker()

	l := log.New(os.Stdout, "winston-chat", log.LUTC)
	
	mux := http.NewServeMux()

	mux.Handle("/worker", handlers.NewWorker(l))
	http.ListenAndServe(":8123", mux)
}
