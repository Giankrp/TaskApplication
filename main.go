package main

import (
	"net/http"

	"github.com/Giankrp/chiBack/db"
	"github.com/Giankrp/chiBack/models"
	"github.com/Giankrp/chiBack/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db.DBConnection()

	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Task{})

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", routes.HomeHandler)
	r.Post("/", routes.PostTaskHandler)
	r.Post("/users", routes.CreateUserHandler)
	r.Get("/users", routes.UsersHandler)
	r.Get("/users/{ID}", routes.GetUserHandler)
	r.Delete("/users/{ID}", routes.DeleteUserHanlder)

	http.ListenAndServe(":8000", r)
}
