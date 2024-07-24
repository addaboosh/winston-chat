package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/addaboosh/winston-chat/config"
	"github.com/addaboosh/winston-chat/store"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	cfg config.Configuration
	store store.Interface
	router *chi.Mux
	l *log.Logger
}

func NewServer(cfg config.Configuration, store store.Interface, l *log.Logger) *Server {
	srv := &Server{
		cfg: cfg,
		store: store,
		router: chi.NewRouter(),
		l: l,
	}	

	srv.routes()

	return srv

}

func (s *Server) Start(ctx context.Context){
	s.l.Printf("Starting server on port %d", s.cfg.Port)
	server := http.Server{
		Addr: fmt.Sprintf(":%d", s.cfg.HTTPServer.Port),
		Handler: s.router,
		IdleTimeout: s.cfg.HTTPServer.IdleTimeout,
		ReadTimeout: s.cfg.HTTPServer.ReadTimeout,
		WriteTimeout: s.cfg.HTTPServer.WriteTimeout,
	}

	shutdownComplete := handleShutdown(func(){
		if err := server.Shutdown(ctx); err != nil {
			s.l.Printf("server.Shutdown failed: %v\n", err)
		}

	})
	
	if err := server.ListenAndServe(); err == http.ErrServerClosed {
		<- shutdownComplete
	} else {
		s.l.Printf("http.ListenAndServe failed: %v", err)
	}
	s.l.Printf("shutdown gracefully")
}

func handleShutdown(onShutdownSignal func()) <- chan struct{} {
	shutdown := make(chan struct{})

	go func(){
		shutdownSignal := make(chan os.Signal, 1)
		signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

		<- shutdownSignal

		onShutdownSignal()
		close(shutdown)
	}()
	return shutdown
}
