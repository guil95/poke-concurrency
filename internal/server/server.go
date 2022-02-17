package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/guil95/poke-concurrency/internal"
	"net/http"
	"strconv"
)

type Server struct {
	router *mux.Router
	uc *internal.UseCase
}

func NewServer(r *mux.Router, uc *internal.UseCase) *Server {
	return &Server{
		router: r,
		uc: uc,
	}
}

func (s *Server) Serve() {
	s.router.HandleFunc("/pokemons/{total:[0-9]+}/{method:[a-z]+}", s.listPokemons).Methods("GET")
}

func (s *Server) listPokemons(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	totalInt, err := strconv.Atoi(params["total"])
	if err != nil {
		return
	}

	pokemons, _ := s.uc.GetPokemons(totalInt, params["method"])

	json.NewEncoder(w).Encode(pokemons)
}
