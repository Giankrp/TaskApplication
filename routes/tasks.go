package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Giankrp/chiBack/db"
	"github.com/Giankrp/chiBack/models"
	"github.com/go-chi/chi/v5"
)

func GetTasksHander(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	db.DB.Find(&tasks)

	json.NewEncoder(w).Encode(&tasks)
}

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	json.NewDecoder(r.Body).Decode(&task)

	DbTask := db.DB.Create(&task)

	if err := DbTask.Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&task)
}

func GetTaskHadler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := chi.URLParam(r, "ID")

	db.DB.First(&task, params)

	if task.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Task not found"))
		return
	}

	json.NewEncoder(w).Encode(&task)
}
