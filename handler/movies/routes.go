package movies

import (
	"github.com/gin-gonic/gin"
)

func (h *MovieHandler) Route(app *gin.Engine) {
	v1 := app.Group("api/v1")
	v1.Use()
	{
		v1.GET("/health", h.HealthGET)
		v1.GET("/search", h.Search)
		v1.GET("/detail", h.MovieDetailByID)
	}
}
