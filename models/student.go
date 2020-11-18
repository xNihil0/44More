package models

import (
	"errors"
	"gorm.io/gorm"
	db "mtdn.io/44More/database"
)

type StudentModel struct {
	gorm.Model `json:"-"`
	StuNo      int64  `json:"stuNo"`
	Name       string `json:"name"`
	Building   string `json:"building"`
	Room       int64  `json:"room"`
	Bed        int64  `json:"bed"`
}

func (s *StudentModel) CreateStudent() (id int64, err error) {
	var student StudentModel
	if err := db.StudentDB.Where("stu_no=?", s.StuNo).First(&student).Error; err == nil {
		return 0, errors.New("existing student number")
	} else if err == gorm.ErrRecordNotFound {
		db.StudentDB.Create(s)
		return s.StuNo, nil
	} else {
		return 0, err
	}
}

func GetStudent(stuno int64) (*StudentModel, error) {
	var student StudentModel
	if err := db.StudentDB.Where("stu_no=?", stuno).First(&student).Error; err != nil {
		return nil, err
	} else {
		return &student, nil
	}
}

func (s *StudentModel) UpdateStudent() (int64, error) {
	var student StudentModel
	if err := db.StudentDB.Where("stu_no=?", s.StuNo).First(&student).Error; err == nil {
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
		db.StudentDB.Save(&student)
		return student.StuNo, nil
	} else {
		return 0, err
	}
}

func RemoveStudent(stuno int64) (int64, error) {
	var student StudentModel
	if err := db.StudentDB.Where("stu_no=?", stuno).First(&student).Error; err == nil {
		db.StudentDB.Where("stu_no=?", stuno).Delete(&StudentModel{})
		return student.StuNo, nil
	} else {
		return 0, err
	}
}
