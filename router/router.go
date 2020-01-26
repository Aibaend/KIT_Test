package router

import (
	"github.com/gorilla/mux"
	"KIT_Test/controller"
)

func SetupRoutes(router *mux.Router)() {
	//dictionary
	router.HandleFunc("/dictionary/create",controller.CreateDictionary).Methods("POST")
	router.HandleFunc("/dictionary/get",controller.GetDictionary).Methods("GET")
}