package database

import (
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var StudentDB *gorm.DB

func init() {
	var err error

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)

	StudentDB, err = gorm.Open(sqlite.Open("44More.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		//log.Panic(err.Error())
		panic(err)
	}
	err = StudentDB.AutoMigrate(&StudentModel{})
	if err != nil {
		panic(err)
	}
}

//
//
//

type StudentModel struct {
	gorm.Model
	StuNo    int64 `gorm:"index;unique"` //todo: string
	Name     string
	Building string
	Room     int64
	Bed      int64
}

func (s *StudentModel) CreateStudent() (id int64, err error) {
	var student StudentModel
	if err := StudentDB.Where("stu_no=?", s.StuNo).First(&student).Error; err == nil {
		return 0, errors.New("existing student number")
	} else if err == gorm.ErrRecordNotFound {
		StudentDB.Create(s)
		return s.StuNo, nil
	} else {
		return 0, err
	}
}

func GetStudent(stuno int64) (*StudentModel, error) {
	var student StudentModel
	if err := StudentDB.Where("stu_no=?", stuno).First(&student).Error; err != nil {
		return nil, err
	} else {
		return &student, nil
	}
}

func (s *StudentModel) UpdateStudent() (int64, error) {
	var student StudentModel
	if err := StudentDB.Where("stu_no=?", s.StuNo).First(&student).Error; err == nil {
		if s.Name != "" {
			student.Name = s.Name
		}
		if s.Building != "" {
			student.Building = s.Building
		}
		if s.Room != 0 {
			student.Room = s.Room
		}
		if s.Bed != 0 {
			student.Bed = s.Bed
		}
		StudentDB.Save(&student)
		return student.StuNo, nil
	} else {
		return 0, err
	}
}

func RemoveStudent(stuno int64) (int64, error) {
	var student StudentModel
	if err := StudentDB.Where("stu_no=?", stuno).First(&student).Error; err == nil {
		StudentDB.Where("stu_no=?", stuno).Delete(&StudentModel{})
		return student.StuNo, nil
	} else {
		return 0, err
	}
}

func (s *StudentModel) GetRoommate() *[]StudentModel {
	var roommates []StudentModel
	StudentDB.Where("building=? AND room=?", s.Building, s.Room).Find(&roommates)
	return &roommates
}
