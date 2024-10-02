package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/addaboosh/winston-chat/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type workerResponse struct {
	Id          uuid.UUID                       `json:"id"`
	Name        string                          `json:"name"`
	Connections map[uuid.UUID]store.IConnection `json:"connections"`
}

func (hr workerResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewWorkerResponse(w store.Worker) workerResponse {
	return workerResponse{
		Id:          w.Id,
		Name:        w.Name,
		Connections: w.Connections,
	}
}
// HTTP POST - Create New Worker

type CreateWorkerRequest struct {
	Name string `json:"name"`
}

func (wr *CreateWorkerRequest) Bind(r *http.Request) error {
	return nil
}

func (s *Server) handleCreateWorker(w http.ResponseWriter, r *http.Request) {

	data := &CreateWorkerRequest{}
	if err := render.Bind(r, data); err != nil {
		fmt.Printf("err: %v data: %v\n", err, data)
		render.Render(w, r, ErrInternalServerError)
		return
	}

	createWorkerParams := store.CreateWorkerParams{
		Name: data.Name,
	}

	wk, err := s.store.Create(createWorkerParams)

	if err != nil {
		render.Render(w, r, ErrInternalServerError)
		return
	}

	w.WriteHeader(201)
	w.Write(nil)
	render.Render(w, r, NewWorkerResponse(wk))

}
// HTTP GET - Get Worker

func (s *Server) handleGetWorker(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "workerId")
	id, err := uuid.Parse(idParam)
	if err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	worker, err := s.store.GetByID(id)

	if err != nil {
		var rnfErr *store.ResourceNotFoundError
		if errors.As(err, &rnfErr) {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrInternalServerError)
		}
		return
	}
	wr := NewWorkerResponse(worker)
	render.Render(w, r, wr)
}

// HTTP GET - Get All Workers

func NewWorkerListResponse(workers []store.Worker) []render.Renderer {
	list := []render.Renderer{}
	for _, worker := range workers {
		wr := NewWorkerResponse(worker)
		list = append(list, wr)
	}
	return list
}

func (s *Server) handleListWorkers(w http.ResponseWriter, r *http.Request) {
	workers, err := s.store.GetAll()
	if err != nil {
		render.Render(w, r, ErrInternalServerError)
		return
	}
	render.RenderList(w, r, NewWorkerListResponse(workers))
}



// HTTP PUT - Update Worker Name

type setNameRequest struct {
	Name string `json:"name"`
}

func (wr *setNameRequest) Bind(r *http.Request) error {
	return nil
}

func (s *Server) handleSetWorkerName(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "workerId")
	id, err := uuid.Parse(idParam)
	if err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	data := &setNameRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	setNameParams := store.SetWorkerNameParams{
		Name: data.Name,
	}

	err = s.store.SetName(id, setNameParams)

	if err != nil {
		var rnfErr *store.ResourceNotFoundError
		if errors.As(err, &rnfErr) {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrInternalServerError)
		}
		return
	}

	w.WriteHeader(204)
	w.Write(nil)
}

// HTTP DELETE - Delete single worker by ID

func (s *Server) handleDeleteWorker(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "workerId")
	id, err := uuid.Parse(idParam)
	if err != nil {
		render.Render(w, r, ErrInternalServerError)
		return
	}

	err = s.store.Delete(id)
	if err != nil {
		var rnfErr *store.ResourceNotFoundError
		if errors.As(err, &rnfErr) {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrInternalServerError)
		}
		return
	}

	w.WriteHeader(200)
	w.Write(nil)

}

// WIP HELPER - HTTP DELETE - Delete ALL Workers

func (s *Server) handleDeleteWorkers(w http.ResponseWriter, r *http.Request) {
	data, err := s.store.GetAll()
	if err != nil {
		render.Render(w, r, ErrInternalServerError)
		return
	}
	for _, w := range data "wss://eventsub.wss.twitch.tv/ws"{
		s.store.Delete(w.Id)
	}
	return
}
