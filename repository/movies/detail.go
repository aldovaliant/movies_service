package movies

import (
	"encoding/json"
	"fmt"
	"movies/entities/model"
)

func (m *movieRepository) MovieDetailByID(id string) (*model.MovieResult, error) {
	resp, err := m.Omdb.requestAPI("id", id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	r := new(model.MovieResult)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return r, nil
}
