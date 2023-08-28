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
	// NOTE: DB CONFIG
	db.DBConnection()

	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Task{})

	// NOTE: ROUTER

	r := chi.NewRouter()

	// NOTE: MIDDLEWARE
	r.Use(middleware.Logger)

	// NOTE: USERS ROUTES
	r.Get("/", routes.HomeHandler)
	r.Post("/users", routes.CreateUserHandler)
	r.Get("/users", routes.UsersHandler)
	r.Get("/users/{ID}", routes.GetUserHandler)
	r.Delete("/users/{ID}", routes.DeleteUserHanlder)

	// NOTE: TASKS ROUTES
	r.Post("/", routes.PostTaskHandler)
	r.Get("/tasks", routes.GetTasksHander)
	r.Get("/tasks/{ID}", routes.GetTaskHadler)
  r.Delete("/tasks/{ID}", routes.DeleteTaskHandler)

	http.ListenAndServe(":8000", r)
}
