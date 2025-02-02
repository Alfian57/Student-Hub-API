package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) internalServerError(c *gin.Context, err error) {
	app.logger.Errorw("internal server error", "method", c.Request.Method, "path", c.Request.URL, "error", err.Error())
	app.writeResponseMessage(c, http.StatusInternalServerError, err.Error())
}

func (app *application) badRequestError(c *gin.Context, err error) {
	app.logger.Infow("bad request error", "method", c.Request.Method, "path", c.Request.URL, "error", err.Error())
	app.writeResponseMessage(c, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundError(c *gin.Context, err error) {
	app.logger.Infow("not found error", "method", c.Request.Method, "path", c.Request.URL, "error", err.Error())
	app.writeResponseMessage(c, http.StatusNotFound, err.Error())
}

func (app *application) unprocessableEntityError(c *gin.Context, err error) {
	app.logger.Infow("unprocessable entity error", "method", c.Request.Method, "path", c.Request.URL, "error", err.Error())
	app.writeResponseMessage(c, http.StatusUnprocessableEntity, err.Error())
}
