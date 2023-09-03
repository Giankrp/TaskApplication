package main

import (
	"net/http"

	"github.com/Giankrp/chiBack/db"
	"github.com/Giankrp/chiBack/models"
	"github.com/Giankrp/chiBack/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
    AllowedOrigins: []string{"https://taskapplication-production.up.railway.app", "http://localhost:5173"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
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
