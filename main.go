package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rohanhonnakatti/sqlite-basic-crud/controllers"
	"github.com/rohanhonnakatti/sqlite-basic-crud/routes"
)

func main() {

	controllers.Init()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	routes.RegisterUserRoutes(r)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("server is up & running..."))
	})

	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
