package main

import (
	"github/Alfian57/student-hub-api/internal/store"

	"github.com/gin-gonic/gin"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	appUrl string
	addr   string
	db     dbConfig
	env    string
}

type dbConfig struct {
	addr string
}

func (app *application) mount() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", app.healthCheckHandler)

		admin := v1.Group("/admin")
		{
			blogCategory := admin.Group("/blog-categories")
			{
				blogCategory.GET("/", app.getBlogCategoryListHandler)
				blogCategory.POST("/", app.createBlogCategoryHandler)
				blogCategory.PUT("/:slug", app.updateBlogCategoryHandler)
				blogCategory.DELETE("/:slug", app.deleteBlogCategoryHandler)
			}

			projectCategory := admin.Group("/project-categories")
			{
				projectCategory.GET("/", app.getProjectCategoryListHandler)
				projectCategory.POST("/", app.createProjectCategoryHandler)
				projectCategory.PUT("/:slug", app.updateProjectCategoryHandler)
				projectCategory.DELETE("/:slug", app.deleteProjectCategoryHandler)
			}
		}
	}

	return r
}

func (app *application) run(r *gin.Engine) error {
	return r.Run(app.config.addr)
}
