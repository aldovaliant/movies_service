package movies

import (
	"movies/entities/model"
	repository "movies/repository/movies"
)

type (
	MovieUsecase interface {
		Search(model.QueryData) (*model.SearchResponse, error)
		MovieDetailByID(string) (*model.MovieResult, error)
	}

	movieUsecase struct {
		movieRepository repository.MovieRepository
	}
)

func NewMovieUsecase(movieRepository repository.MovieRepository) MovieUsecase {
	return &movieUsecase{
		movieRepository: movieRepository,
	}
}
