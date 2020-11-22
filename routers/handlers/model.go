package handlers

import (
	"mtdn.io/44More/database"
	"time"
)

// database
//type StudentModel struct {
//	gorm.Model
//	StuNo    int64 `gorm:"index;unique"` //todo: string
//	Name     string
//	Building string
//	Room     int64
//	Bed      int64
//}

// router api
type InputStudent struct {
	//gorm.Model `json:"-"`
	StuNo    int64  `json:"stuNo"`
	Name     string `json:"name"`
	Building string `json:"building"`
	Room     int64  `json:"room"`
	Bed      int64  `json:"bed"`
}

type OutputStudent struct {
	//gorm.Model `json:"-"`
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	StuNo    int64  `json:"stuNo"`
	Name     string `json:"name"`
	Building string `json:"building"`
	Room     int64  `json:"room"`
	Bed      int64  `json:"bed"`
}

func ToOutputStudent(studentModel []*database.StudentModel) []*OutputStudent {
	if len(studentModel) == 0 {
		return nil
	}

	outputStudents := make([]*OutputStudent, 0, 100)
	for _, v := range studentModel {
		outputStudents = append(outputStudents, &OutputStudent{
			ID:        v.ID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			StuNo:     v.StuNo,
			Name:      v.Name,
			Building:  v.Building,
			Room:      v.Room,
			Bed:       v.Bed,
		})
	}
	return outputStudents
}
