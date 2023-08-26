package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Giankrp/chiBack/db"
	"github.com/Giankrp/chiBack/models"
)

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
