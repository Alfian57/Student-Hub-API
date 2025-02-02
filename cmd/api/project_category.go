package main

import (
	"errors"
	"github/Alfian57/student-hub-api/internal/store"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

type CreateProjectCategoryPayload struct {
	Name string `json:"name" form:"name" validate:"required,max=100"`
}

type UpdateProjectCategoryPayload struct {
	Name string `json:"name" form:"name" validate:"required,max=100"`
}

func (app *application) getProjectCategoryListHandler(c *gin.Context) {
	var queryParam store.CategoryQueryParam
	if err := queryParam.Parse(c); err != nil {
		app.badRequestError(c, err)
		return
	}

	if err := Validator.Struct(queryParam); err != nil {
		app.badRequestError(c, err)
		return
	}

	categories, err := app.store.Category.GetAllProjectCategory(c.Request.Context(), queryParam)
	if err != nil {
		app.internalServerError(c, err)
		return
	}

	app.writeResponseData(c, http.StatusOK, categories)
}

func (app *application) createProjectCategoryHandler(c *gin.Context) {
	var payload CreateProjectCategoryPayload
	if err := c.ShouldBind(&payload); err != nil {
		app.badRequestError(c, err)
		return
	}

	if err := Validator.Struct(payload); err != nil {
		app.unprocessableEntityError(c, err)
		return
	}

	category := store.Category{
		ID:   uuid.New(),
		Slug: slug.Make(payload.Name),
		Name: payload.Name,
		Type: store.CategoryTypeProject,
	}

	err := app.store.Category.Create(c.Request.Context(), &category)
	if err != nil {
		app.internalServerError(c, err)
		return
	}

	app.writeResponseData(c, http.StatusOK, category)
}

func (app *application) updateProjectCategoryHandler(c *gin.Context) {
	categorySlug := c.Param("slug")

	var payload UpdateProjectCategoryPayload
	if err := c.ShouldBind(&payload); err != nil {
		app.badRequestError(c, err)
	}

	if err := Validator.Struct(payload); err != nil {
		app.unprocessableEntityError(c, err)
	}

	category := store.Category{
		Slug:      slug.Make(payload.Name),
		Name:      payload.Name,
		Type:      store.CategoryTypeProject,
		UpdatedAt: time.Now().UTC(),
	}

	err := app.store.Category.Update(c.Request.Context(), categorySlug, &category)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			app.notFoundError(c, err)
		} else {
			app.internalServerError(c, err)
		}
		return
	}

	app.writeResponseMessage(c, http.StatusOK, "Category successfully updated")
}

func (app *application) deleteProjectCategoryHandler(c *gin.Context) {
	categorySlug := c.Param("slug")

	err := app.store.Category.Delete(c.Request.Context(), categorySlug)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			app.notFoundError(c, err)
		} else {
			app.internalServerError(c, err)
		}
		return
	}

	app.writeResponseMessage(c, http.StatusOK, "Category successfully deleted")
}
