package main

import (
	"context"
	"log"
	"os"

	"github.com/addaboosh/winston-chat/api"
	"github.com/addaboosh/winston-chat/config"
	"github.com/addaboosh/winston-chat/store"
)



func main() {
	logger := log.New(os.Stdout, "winston-chat - ", log.LUTC)
	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Printf("cfg: %v", cfg)
	store := store.NewMemoryWorkerStore()
	server := api.NewServer(cfg, store, logger)

	server.Start(ctx)

}
