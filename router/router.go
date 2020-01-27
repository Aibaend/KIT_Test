package router

import (
	"github.com/gorilla/mux"
	"KIT_Test/controller"
)

func SetupRoutes(router *mux.Router)() {
	//dictionary
	router.HandleFunc("/dictionary/create",controller.CreateDictionary).Methods("POST")
	router.HandleFunc("/dictionary/get",controller.GetDictionary).Methods("GET")
	router.HandleFunc("/dictionary/delete/id:{id}",controller.DeleteDictionary).Methods("DELETE")
	router.HandleFunc("/dictionary/update/id:{id}",controller.UpdateDictionary).Methods("PUT")
	router.HandleFunc("/dictionary/find/{word}",controller.FindDictionaryElement).Methods("GET")
}