package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	. "mtdn.io/44More/models"
	"net/http"
	"strconv"
)

func CreateStudentHandler(c *gin.Context) {
	stuno, _ := strconv.ParseInt(c.PostForm("stuno"), 10, 64)
	name := c.PostForm("name")
	building := c.PostForm("building")
	room, _ := strconv.ParseInt(c.PostForm("room"), 10, 64)
	bed, _ := strconv.ParseInt(c.PostForm("bed"), 10, 64)
	var status int
	var msg string
	errorCount := 0
	var student *StudentModel
	if name == "" || building == "" || room == 0 || bed == 0 {
		status = http.StatusBadRequest
		msg = "not enough parameters"
		errorCount++
	} else {
		student = &StudentModel{
			StuNo:    stuno,
			Name:     name,
			Building: building,
			Room:     room,
			Bed:      bed,
		}
		_, err := student.CreateStudent()
		if err == nil {
			status = http.StatusOK
			msg = "ok"
		} else {
			status = http.StatusInternalServerError
			msg = err.Error()
			errorCount++
		}
	}

	c.JSON(status, gin.H{
		"msg":    msg,
		"data":   student,
		"error":  errorCount,
	})
}

func GetStudentHandler(c *gin.Context) {
	stuno, _ := strconv.ParseInt(c.Param("stuno"), 10, 64)
	student, err := GetStudent(stuno)
	var status int
	var msg string
	errorCount := 0
	if err == nil {
		status = http.StatusOK
		msg = "ok"
	} else if err == gorm.ErrRecordNotFound {
		status = http.StatusNotFound
		msg = err.Error()
		errorCount++
	} else {
		status = http.StatusInternalServerError
		msg = err.Error()
		errorCount++
	}

	c.JSON(status, gin.H{
		"msg":    msg,
		"data":   student,
		"error":  errorCount,
	})
}

func UpdateStudentHandler(c *gin.Context) {
	stuno, _ := strconv.ParseInt(c.PostForm("stuno"), 10, 64)
	name := c.PostForm("name")
	building := c.PostForm("building")
	room, _ := strconv.ParseInt(c.PostForm("room"), 10, 64)
	bed, _ := strconv.ParseInt(c.PostForm("bed"), 10, 64)

	student := &StudentModel{
		StuNo:    stuno,
		Name:     name,
		Building: building,
		Room:     room,
		Bed:      bed,
	}
	var status int
	var msg string
	errorCount := 0
	_, err := student.UpdateStudent()
	if err == nil {
		status = http.StatusOK
		msg = "ok"
	} else {
		status = http.StatusInternalServerError
		msg = err.Error()
		errorCount++
	}

	c.JSON(status, gin.H{
		"msg":    msg,
		"data":   student,
		"error":  errorCount,
	})
}

func RemoveStudentHandler(c *gin.Context) {
	stuno, _ := strconv.ParseInt(c.Param("stuno"), 10, 64)
	_, err := RemoveStudent(stuno)
	var status int
	var msg string
	errorCount := 0
	if err == nil {
		status = http.StatusOK
		msg = "ok"
	} else if err == gorm.ErrRecordNotFound {
		status = http.StatusNotFound
		msg = err.Error()
		errorCount++
	} else {
		status = http.StatusInternalServerError
		msg = err.Error()
		errorCount++
	}

	c.JSON(status, gin.H{
		"msg":    msg,
		"data":   stuno,
		"error":  errorCount,
	})
}

func GetRoommateHandler(c *gin.Context) {
	stuno, _ := strconv.ParseInt(c.Param("stuno"), 10, 64)
	student, err := GetStudent(stuno)
	var status int
	var msg string
	errorCount := 0
	var roommates *[]StudentModel
	if err == nil {
		status = http.StatusOK
		msg = "ok"
		roommates = student.GetRoommate()
	} else if err == gorm.ErrRecordNotFound {
		status = http.StatusNotFound
		msg = err.Error()
		errorCount++
	} else {
		status = http.StatusInternalServerError
		msg = err.Error()
		errorCount++
	}

	c.JSON(status, gin.H{
		"msg":    msg,
		"data":   roommates,
		"error":  errorCount,
	})
}
