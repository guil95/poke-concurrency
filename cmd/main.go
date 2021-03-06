package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guil95/poke-concurrency/internal"
	"github.com/guil95/poke-concurrency/internal/client"
	"github.com/guil95/poke-concurrency/internal/server"
)

func main()  {
	r := mux.NewRouter()
	uc := internal.NewUseCase(client.NewClient())
	s := server.NewServer(r, uc)
	s.Serve()
	log.Println("Running")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		return
	}
}
