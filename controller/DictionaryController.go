package controller

import (
	"KIT_Test/model"
	u "KIT_Test/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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

var DeleteDictionary = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err !=nil{
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	resp := model.DeleteDictionary(uint(id))
	u.Respond(w, resp)
}


var UpdateDictionary  = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	dicti := &model.Dictionary{}
	err1 := json.NewDecoder(r.Body).Decode(dicti)
	if err1 != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	data := dicti.UpdateDictionary(uint(id))

	resp := u.Message(true,"success")
	resp["data"] = data
	u.Respond(w, resp)
}

var FindDictionaryElement  = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	keyword, _ := params["word"]
	if keyword ==""{
		u.Respond(w,u.Message(false,"Keyword is empty"))
		return
	}
	data := model.FindDictionary(keyword)
	if len(data) == 0{
		resp := u.Message(false,"Data is not found")
		u.Respond(w,resp)
	}else {
		resp := u.Message(true, "success")
		resp["data"] = data
		u.Respond(w, resp)
	}
}