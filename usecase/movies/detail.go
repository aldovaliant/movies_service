package movies

import (
	"fmt"
	"movies/entities/model"
)

func (m *movieUsecase) MovieDetailByID(id string) (*model.MovieResult, error) {
	res, err := m.movieRepository.MovieDetailByID(id)
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	return res, nil
}
