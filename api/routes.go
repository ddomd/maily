package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)


func (s *Server) registerRoutes() {
	corsOpts := cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}
	
	router := chi.NewRouter()
	restRouter := chi.NewRouter()
	
	router.Use(cors.Handler(corsOpts))
	
	router.Mount("/rest", restRouter)
	
	restRouter.Get("/batch/{limit}&{offset}", s.HandleGetBatchEmail)
	restRouter.Get("/get", s.HandleGetEmail)
	
	restRouter.Post("/create", s.HandleCreateEmail)
	restRouter.Put("/update", s.HandleUpdateEmail)
	restRouter.Delete("/delete", s.HandleDeleteEmail)

	s.Router = router
}