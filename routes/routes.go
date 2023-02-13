package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/namelew/personality-api/controllers"
	"github.com/namelew/personality-api/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentType)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.GetPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriarPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletarPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditarPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
