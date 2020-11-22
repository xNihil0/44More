package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	//student := router.Group("/student")
	{
		//student.POST("/create", RequestEntry(CreateStudent))
		//student.GET("/info/:stuno", RequestEntry(GetStudentHandler))
		//student.PUT("/update", RequestEntry(UpdateStudentHandler))
		//student.DELETE("/remove/:stuno", RequestEntry(RemoveStudentHandler))
		//student.GET("/roommate/:stuno", RequestEntry(GetRoommateHandler))
	}

	return router
}

func MakeSuccessReturn(data interface{}) (status int, responseBody interface{}) {
	return http.StatusOK, SuccessReturn{
		Msg:   "success",
		Data:  data,
		Error: 0,
	}
}

func MakeErrorReturn(status int, code int, msg string) (int, interface{}) {
	return status, ErrorReturn{
		Msg:  msg,
		Code: code,
	}
}

type SuccessReturn struct {
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Error int         `json:"error"`
}

type ErrorReturn struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

func RequestEntry(f func(ctx *gin.Context) (int, interface{})) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(f(c))
	}
}
