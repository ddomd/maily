package api

import (
	"log"
	"net/http"

	"github.com/ddomd/maily/internal/mdb"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	Port   string
	Router *chi.Mux
	DB     *mdb.MDB
}

func NewServer(port string, db *mdb.MDB) *Server {
	server := Server{
		Port: port,
		DB: db,
	}
	server.registerRoutes()
	return &server
}

func (s *Server) Serve() {
	log.Fatal(http.ListenAndServe(s.Port, s.Router))
}