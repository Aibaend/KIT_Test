package controller

import (
	"KIT_Test/model"
	u "KIT_Test/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var CreateDictionary = func(w http.ResponseWriter, r *http.Request) {
	dicti := &model.Dictionary{}
	err := json.NewDecoder(r.Body).Decode(dicti) //decode the request body into struct and failed if any error occur
	fmt.Print(err)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := dicti.Create() //Create account
	u.Respond(w, resp)
}

var GetDictionary = func(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	id := v.Get("id")
	var result  int
	if id !="" {
		scorm, err := strconv.Atoi(id)
		if err != nil {
			//The passed path parameter is not an integer
			u.Respond(w, u.Message(false, "There was an error in your request"))
			return
		}
		result = scorm
	}else {
		result = 0
	}


	data := model.GetDictionary(uint(result))

	u.Respond(w, data)
}