package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"

	"github.com/tupis/microsservice-go-chi-redis/handler"
)

func loadRoutes() *chi.Mux {

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	router.Route("/orders", loadOrdersRoutes)

	return router
}

func loadOrdersRoutes(router chi.Router) {
	orderHandler := &handler.Order{}

	router.Get("/list", orderHandler.List)
	router.Post("/create", orderHandler.Create)
	router.Get("/{id}", orderHandler.GetByID)
	router.Put("/{id}", orderHandler.UpdateById)
	router.Delete("/{id}", orderHandler.DeleteById)
}
