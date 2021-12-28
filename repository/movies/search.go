package movies

import (
	"encoding/json"
	"fmt"
	"movies/entities/model"
)

func (m *movieRepository) Search(query model.QueryData) (*model.SearchResponse, error) {
	resp, err := m.Omdb.requestAPI("search", query.Title, query.Year, query.SearchType, query.Page)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	r := new(model.SearchResponse)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// if r.Response == "False" {
	// 	return r, errors.New(r.Error)
	// }

	return r, nil
}
