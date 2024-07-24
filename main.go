package main

import (
	"context"
	"log"

	"github.com/addaboosh/winston-chat/api"
	"github.com/addaboosh/winston-chat/config"
	"github.com/addaboosh/winston-chat/store"
)



func main() {
	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	store := store.NewMemoryWorkerStore()
	server := api.NewServer(cfg, store)
	server.Start(ctx)
}
