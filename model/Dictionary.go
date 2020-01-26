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

func DeleteDictionary(id uint) (map[string] 	interface{}) {
	web := &Dictionary{}

	err:=GetDB().Table("dictionaries").Where("id = ?", id).First(web).Error
	if err!=nil{
		if err == gorm.ErrRecordNotFound{
			return u.Message(false,"id is not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}


	db.Where("id = ?", id).Delete(&Dictionary{})
	response := u.Message(true, "Dictionary has been deleted")
	return response
}

func (dicti *Dictionary) UpdateDictionary(id uint) (*Dictionary) {

	temp_dicti :=&Dictionary{}
	err := GetDB().Table("dictionaries").Where("id = ?", id).First(temp_dicti).Error
	if err != nil {
		return nil
	}

	temp_dicti = dicti

	db.Model(&temp_dicti).Where("id = ?", id).Update(dicti)


	return temp_dicti

}
