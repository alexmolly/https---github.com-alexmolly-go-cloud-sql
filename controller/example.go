package controller

import (
	"net/http"

	"github.com/alexmolly/gomvcboilerplate/services"
	"github.com/gin-gonic/gin"
)

type ExampleController struct{}

func (controller ExampleController) ExampleHandler(c *gin.Context) {
	var responseString = services.ExampleService.SayHi()
	var response = SuccessResponse{
		Status: true,
		Msg:    responseString,
	}
	c.AbortWithStatusJSON(http.StatusOK, response)
}

func (controller ExampleController) CheckDBConnection(c *gin.Context) {
	var responseString = "Connected to DB"
	var sqlDB, _ = services.ExampleService.DB.DB()
	if err := sqlDB.Ping(); err != nil {
		responseString = "Cannot connect to DB"
	}
	var response = SuccessResponse{
		Status: true,
		Msg:    responseString,
	}
	c.AbortWithStatusJSON(http.StatusOK, response)
}
