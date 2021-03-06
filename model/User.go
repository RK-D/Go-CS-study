package model

import "github.com/jinzhu/gorm"

//定义model
type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	//grom:不要加空格
	Telephone string `gorm:"varchar(110;not null;unique"`
	Password  string `gorm:"size:255;not null"`
}
