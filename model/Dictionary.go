package model

import (
	u "KIT_Test/utils"
	"github.com/jinzhu/gorm"
)

type Dictionary struct {
	gorm.Model
	Name 			string
	Description 	string
	Dictionary 		[]*Dictionary `json:",omitempty"`

}


func (dicti *Dictionary) Create() (map[string] interface{}) {
	if len(dicti.Dictionary)>0{
		for i  := range dicti.Dictionary{
			GetDB().Create(dicti.Dictionary[i])
		}
	}else {
		GetDB().Create(dicti)
	}
	response := u.Message(true, "Dictionary has been created")
	return response
}


func GetDictionary(id uint) (map[string] interface{})  {
	dicti := [] Dictionary{}
	if id ==0{
		db.Find(&dicti)

	}else {
		db.First(&dicti,id)

	}
	resp := u.Message(true,"Dictionary found")
	resp["data"] = dicti
	return resp

}