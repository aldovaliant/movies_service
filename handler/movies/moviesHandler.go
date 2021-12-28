package movies

import (
	"movies/entities/model"
	"movies/usecase/movies"
	"movies/utils"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	movieUsecase movies.MovieUsecase
}

func NewMovieHandler(movieUsecase movies.MovieUsecase) *MovieHandler {
	return &MovieHandler{
		movieUsecase: movieUsecase,
	}
}

func (h *MovieHandler) Search(c *gin.Context) {
	query := model.QueryData{}
	query.Title, _ = c.GetQuery("searchword")
	query.Year, _ = c.GetQuery("year")
	query.SearchType, _ = c.GetQuery("type")
	query.Page, _ = c.GetQuery("pagination")

	res, err := h.movieUsecase.Search(query)
	if err != nil {
		if err.Error() == "data not found" {
			utils.Response(c, 404, err.Error(), nil)
			return
		}
		utils.Response(c, 500, "something went wrong", err)
		return
	}
	if res.Response == "False" {
		utils.Response(c, 400, "Bad Request", res)
	} else {
		utils.Response(c, 200, "Success", res)
	}
}

func (h *MovieHandler) MovieDetailByID(c *gin.Context) {
	id, _ := c.GetQuery("id")
	res, err := h.movieUsecase.MovieDetailByID(id)
	if err != nil {
		if err.Error() == "data not found" {
			utils.Response(c, 404, err.Error(), nil)
			return
		}
		utils.Response(c, 500, "something went wrong", err)
		return
	}
	if res.Response == "False" {
		utils.Response(c, 400, "Bad Request", res)
	} else {
		utils.Response(c, 200, "Success", res)
	}
}

func (h *MovieHandler) HealthGET(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "UP",
	})
}
