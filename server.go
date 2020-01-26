package main

import (
	"KIT_Test/model"
	"github.com/gorilla/mux"
	"github.com/liamylian/jsontime"
	"log"
	"net/http"
	"os"
	"fmt"
	"time"
)



func main()  {
	router := mux.NewRouter()
	model.GetDB()
	SetupRoutes(router)
	router.Use(mux.CORSMethodMiddleware(router))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8078" //localhost
	}
	handler := LogMiddleware(router)
	fmt.Println("Server is running on port 8078")

	err := http.ListenAndServe(":" + port,handler) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}

}
