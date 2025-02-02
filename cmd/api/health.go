package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) healthCheckHandler(c *gin.Context) {
	app.writeResponseMessage(c, http.StatusOK, "ok")
}
