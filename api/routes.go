package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func (s *Server) routes() {
	s.router.Use(render.SetContentType(render.ContentTypeJSON))
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Heartbeat("/health"))

	s.router.Route("/api/workers", func(r chi.Router) {
		r.Get("/", s.handleListWorkers)
		r.Post("/", s.handleCreateWorker)
		r.Delete("/", s.handleDeleteWorkers)
		r.Route("/{workerId}", func(r chi.Router) {
			r.Get("/", s.handleGetWorker)
			r.Put("/", s.handleSetWorkerName)
			r.Delete("/", s.handleDeleteWorker)
		})
	})
}
