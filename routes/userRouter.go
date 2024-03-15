package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/rohanhonnakatti/sqlite-basic-crud/controllers"
)

func RegisterUserRoutes(r chi.Router) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/getusers", controllers.GetUsers)
		r.Get("/getuser/{id}", controllers.GetUser)
		r.Post("/adduser", controllers.CreateUser)
		r.Delete("/deleteuser/{id}", controllers.DeleteUser)
		r.Put("/updateuser/{id}", controllers.UpdateUser)
	})
}
