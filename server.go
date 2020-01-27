package main

import
(
	"KIT_Test/model"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"fmt"
	r "KIT_Test/router"
)



func main()  {
	router := mux.NewRouter()
	model.GetDB()
	r.SetupRoutes(router)
	router.Use(mux.CORSMethodMiddleware(router))

	f, err1 := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err1 != nil {
		log.Println(err1)
	}
	defer f.Close()
	log.SetOutput(f)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8078" //localhost
	}
	//handler := LogMiddleware(router)
	fmt.Println("Server is running on port 8078")
	handler := CorsMiddleware(router)
	err := http.ListenAndServe(":" + port,handler)
	if err != nil {
		fmt.Print(err)
	}

}
