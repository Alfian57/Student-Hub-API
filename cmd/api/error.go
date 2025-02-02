package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) internalServerError(c *gin.Context, err error) {
	log.Printf("internal server error: %s path: %s error %s", c.Request.Method, c.Request.URL, err.Error())
	app.writeResponseMessage(c, http.StatusInternalServerError, err.Error())
}

func (app *application) badRequestError(c *gin.Context, err error) {
	log.Printf("bad request error: %s path: %s error %s", c.Request.Method, c.Request.URL, err.Error())
	app.writeResponseMessage(c, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundError(c *gin.Context, err error) {
	log.Printf("not found error: %s path: %s error %s", c.Request.Method, c.Request.URL, err.Error())
	app.writeResponseMessage(c, http.StatusNotFound, err.Error())
}

func (app *application) unprocessableEntityError(c *gin.Context, err error) {
	log.Printf("unprocessable entity error: %s path: %s error %s", c.Request.Method, c.Request.URL, err.Error())
	app.writeResponseMessage(c, http.StatusUnprocessableEntity, err.Error())
}
