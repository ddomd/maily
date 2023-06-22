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
	
	restRouter.Get("/get/{id}", s.HandleGetEmail)
	restRouter.Get("/all", s.HandleGetAllEmails)
	restRouter.Get("/subs", s.HandleGetAllSubscribed)
	restRouter.Get("/batch_subs/limit={limit}&offset={offset}", s.HandleGetBatchSubscribed)
	restRouter.Get("/batch/limit={limit}&offset={offset}", s.HandleGetBatchEmail)
	
	restRouter.Post("/create", s.HandleCreateEmail)
	restRouter.Put("/update/{id}", s.HandleUpdateEmail)
	restRouter.Delete("/delete/{id}", s.HandleDeleteEmail)
	restRouter.Delete("/delete_unsub", s.HandleDeleteUnsubscribed)
	restRouter.Delete("/delete_before", s.HandleDeleteUnsubscribedBefore)

	s.Router = router
}