package router

import (
	"net/http"
	"server/controllers"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

// handler which routes req for all operations needed to be done
func Routes() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://127.0.0.1:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	
	// @ Routes and executed fucntion on those
	router.Get("/health",func(w http.ResponseWriter,req *http.Request) {
		w.WriteHeader(http.StatusOK)
		_,err := w.Write([]byte("Coffee api is running fine🚀🚀..."))
		if err!= nil {
			return
		}
	})
	router.Get("/api/coffees/all",controllers.GetAllCoffees)
	router.Post("/api/coffees/create",controllers.CreateCoffee)
	router.Get("/api/coffees/id/{id:[0-9a-fA-F-]+}",controllers.GetCoffeeByID)
	router.Get("/api/coffees/name/{name}",controllers.GetCoffeeByName) // get coffee by name
	router.Put("/api/coffees/{id:[0-9a-fA-F-]+}",controllers.UpdateCoffeeByID)
	router.Delete("/api/coffees/{id:[0-9a-fA-F-]+}",controllers.DeleteCoffee)


	// $ returning router which would be used by the handler by main server
	return router
}
