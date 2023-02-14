package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/namelew/personality-api/database"
	"github.com/namelew/personality-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var personalidades []models.Personalidade
	database.DB.Find(&personalidades)
	json.NewEncoder(w).Encode(personalidades)
}

func GetPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade

	database.DB.First(&personalidade, id)

	json.NewEncoder(w).Encode(personalidade)
}

func CriarPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade

	json.NewDecoder(r.Body).Decode(&personalidade)

	go database.DB.Create(&personalidade)

	json.NewEncoder(w).Encode(&personalidade)
}

func DeletarPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade

	go database.DB.Delete(&personalidade, id)

	json.NewEncoder(w).Encode(personalidade)
}

func EditarPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade

	database.DB.First(&personalidade, id)

	json.NewDecoder(r.Body).Decode(&personalidade)

	go database.DB.Save(&personalidade)

	json.NewEncoder(w).Encode(personalidade)
}
