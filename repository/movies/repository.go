package movies

import (
	"movies/entities/model"
)

type (
	MovieRepository interface {
		Search(model.QueryData) (*model.SearchResponse, error)
		MovieDetailByID(string) (*model.MovieResult, error)
	}

	movieRepository struct {
		// Conn *gorm.DB
		Omdb OmdbApi
	}
)

func NewMovieRepository() MovieRepository {
	return &movieRepository{
		// Conn: dbCon,
		Omdb: *Init("faf7e5bb"),
	}
}
