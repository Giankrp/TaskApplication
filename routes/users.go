package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Giankrp/chiBack/db"
	"github.com/Giankrp/chiBack/models"
	"github.com/go-chi/chi/v5"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	db.DB.Find(&users)
	for i := range users {
		db.DB.Model(&users[i]).Association("Tasks").Find(&users[i].Tasks)
	}
	json.NewEncoder(w).Encode(&users)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)
	DbUser := db.DB.Create(&user)

	if err := DbUser.Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	params := chi.URLParam(r, "ID")
	db.DB.First(&user, params)

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHanlder(w http.ResponseWriter, r *http.Request) {
	var user models.User

	params := chi.URLParam(r, "ID")

	db.DB.First(&user, params)

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	w.WriteHeader(http.StatusOK)
	db.DB.Unscoped().Delete(&user)
}
