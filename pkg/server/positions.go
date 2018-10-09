package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nii236/margin/pkg/positions"
)

// Positions holds the handlers
type Positions struct {
	Repo positions.Repo
}

// NewPositions returns a new router for positions
func NewPositions(Repo positions.Repo) chi.Router {
	r := chi.NewRouter()

	ctrl := &Positions{
		Repo: Repo,
	}
	r.Post("/list", ctrl.List)
	r.Post("/open", ctrl.Open)
	r.Post("/close", ctrl.Close)

	return r
}

// List handles the HTTP List requests
func (ctrl *Positions) List(w http.ResponseWriter, r *http.Request) {
	list, err := ctrl.Repo.List()
	if err != nil {
		fmt.Println(err)
		return
	}
	resp := &positions.ListResponse{
		Positions: list,
	}
	json.NewEncoder(w).Encode(resp)
}

// Open handles the HTTP Open requests
func (ctrl *Positions) Open(w http.ResponseWriter, r *http.Request) {
	resp := &positions.OpenResponse{}
	json.NewEncoder(w).Encode(resp)
}

// Close handles the HTTP Close requests
func (ctrl *Positions) Close(w http.ResponseWriter, r *http.Request) {
	resp := &positions.CloseResponse{}
	json.NewEncoder(w).Encode(resp)
}
