package main

import (
	"errors"
	"github/Alfian57/student-hub-api/internal/store"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BlockBlogPayload struct {
	Reason string `json:"reason" form:"reason" validate:"min=10,max=65535"`
}

func (app *application) getBlogListHandler(c *gin.Context) {
	var queryParam store.BlogQueryParam

	if err := queryParam.Parse(c); err != nil {
		app.internalServerError(c, err)
		return
	}

	if err := Validator.Struct(queryParam); err != nil {
		app.badRequestError(c, err)
		return
	}

	log.Println("log", queryParam.Sort)

	blogs, err := app.store.Blog.GetAll(c.Request.Context(), queryParam)
	if err != nil {
		app.internalServerError(c, err)
		return
	}

	app.writeResponseData(c, http.StatusOK, blogs)
}

func (app *application) getBlogDetailHandler(c *gin.Context) {
	blogSlug := c.Param("slug")

	blog, err := app.store.Blog.GetBySlug(c.Request.Context(), blogSlug)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			app.notFoundError(c, err)
		} else {
			app.internalServerError(c, err)
		}
		return
	}

	app.writeResponseData(c, http.StatusOK, blog)
}

func (app *application) blockBlogHandler(c *gin.Context) {
	blogSlug := c.Param("slug")

	var payload BlockBlogPayload
	if err := c.ShouldBind(&payload); err != nil {
		app.badRequestError(c, err)
		return
	}

	if err := Validator.Struct(payload); err != nil {
		app.unprocessableEntityError(c, err)
		return
	}

	blog, err := app.store.Blog.GetBySlug(c.Request.Context(), blogSlug)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			app.notFoundError(c, err)
		} else {
			app.internalServerError(c, err)
		}
		return
	}

	blockedBlog := store.BlockedBlog{
		ID:     uuid.New(),
		BlogID: blog.ID,
		Reason: payload.Reason,
	}

	if err := app.store.Blog.Block(c.Request.Context(), blockedBlog); err != nil {
		app.internalServerError(c, err)
		return
	}

	app.writeResponseMessage(c, http.StatusOK, "Blog successfully blocked")
}
