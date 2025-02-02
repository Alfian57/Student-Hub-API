package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init() {
	Validator = validator.New(validator.WithRequiredStructEnabled())
}

func (app *application) writeResponseMessage(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"message": message,
	})
}

func (app *application) writeResponseData(c *gin.Context, status int, data any) {
	c.JSON(status, gin.H{
		"data": data,
	})
}
