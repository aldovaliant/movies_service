package movies

import (
	"fmt"
	"movies/entities/model"
)

func (m *movieUsecase) Search(query model.QueryData) (*model.SearchResponse, error) {
	res, err := m.movieRepository.Search(query)
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	return res, nil
}
