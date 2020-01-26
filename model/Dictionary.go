package model

import "github.com/jinzhu/gorm"

type Dictionary struct {
	gorm.Model
	Name 			string
	Description 	string
}