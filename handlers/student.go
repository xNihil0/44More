package handlers

import (
	"github.com/gin-gonic/gin"
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
	var msg string
	var errorList string
	var student *StudentModel
	errorCount := 0
	if name == "" || building == "" || room == 0 || bed == 0 {
		msg = "failed"
		errorList += "not enough parameters\n"
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
			msg = "ok"
		} else {
			msg = "failed"
			errorList += err.Error() + "\n"
			errorCount++
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":    msg,
		"data":   student,
		"error":  errorCount,
		"errors": errorList,
	})
}

func GetStudentHandler(c *gin.Context) {
	stuno, _ := strconv.ParseInt(c.Param("stuno"), 10, 64)
	student, err := GetStudent(stuno)
	var msg string
	var errorList string
	errorCount := 0
	if err == nil {
		msg = "ok"
	} else {
		msg = "failed"
		errorList += err.Error() + "\n"
		errorCount++
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":    msg,
		"data":   student,
		"error":  errorCount,
		"errors": errorList,
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
	var msg string
	var errorList string
	errorCount := 0
	_, err := student.UpdateStudent()
	if err == nil {
		msg = "ok"
	} else {
		msg = "failed"
		errorList += err.Error()
		errorCount++
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":    msg,
		"data":   student,
		"error":  errorCount,
		"errors": errorList,
	})
}

func RemoveStudentHandler(c *gin.Context) {
	stuno, _ := strconv.ParseInt(c.Param("stuno"), 10, 64)
	_, err := RemoveStudent(stuno)
	var msg string
	var errorList string
	errorCount := 0
	if err == nil {
		msg = "ok"
	} else {
		msg = "failed"
		errorList += err.Error() + "\n"
		errorCount++
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":    msg,
		"data":   stuno,
		"error":  errorCount,
		"errors": errorList,
	})
}

func GetRoommateHandler(c *gin.Context) {
	stuno, _ := strconv.ParseInt(c.Param("stuno"), 10, 64)
	student, err := GetStudent(stuno)
	var msg string
	var errorList string
	errorCount := 0
	var roommates *[]StudentModel
	if err == nil {
		roommates = student.GetRoommate()
	} else {
		msg = "failed"
		errorList += err.Error() + "\n"
		errorCount++
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":    msg,
		"data":   roommates,
		"error":  errorCount,
		"errors": errorList,
	})
}
