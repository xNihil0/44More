package main

import (
	"github.com/gin-gonic/gin"
	. "mtdn.io/44More/handlers"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	student := router.Group("/student")
	{
		student.POST("/create", CreateStudentHandler)
		student.GET("/info/:stuno", GetStudentHandler)
		student.PUT("/update", UpdateStudentHandler)
		student.DELETE("/remove/:stuno", RemoveStudentHandler)
	}

	return router
}
