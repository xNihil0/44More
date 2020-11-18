package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type StudentModel struct {
	gorm.Model `json:"-"`
	StuNo      int64  `json:"stuNo"`
	Name       string `json:"name"`
	Building   string `json:"building"`
	Room       int64  `json:"room"`
	Bed        int64  `json:"bed"`
}

var StudentDB *gorm.DB

func init() {
	var err error
	StudentDB, err = gorm.Open(sqlite.Open("44More.db"), &gorm.Config{})
	if err != nil {
		log.Panic(err.Error())
	}
	StudentDB.AutoMigrate(&StudentModel{})
}
