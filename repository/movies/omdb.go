package movies

import (
	"errors"
	"net/http"
	"net/url"
)

const (
	baseURL  = "https://www.omdbapi.com/?"
	plot     = "full"
	tomatoes = "true"

	MovieSearch   = "movie"
	SeriesSearch  = "series"
	EpisodeSearch = "episode"
)

type OmdbApi struct {
	apiKey string
}

func Init(apiKey string) *OmdbApi {
	return &OmdbApi{apiKey: apiKey}
}

func (api *OmdbApi) requestAPI(apiCategory string, params ...string) (resp *http.Response, err error) {
	var URL *url.URL
	URL, err = url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	// Checking for invalid category
	if len(params) > 1 && params[2] != "" {
		if params[2] != MovieSearch &&
			params[2] != SeriesSearch &&
			params[2] != EpisodeSearch {
			return nil, errors.New("Invalid search category- " + params[2])
		}
	}
	parameters := url.Values{}
	parameters.Add("apikey", api.apiKey)

	switch apiCategory {
	case "search":
		parameters.Add("s", params[0])
		parameters.Add("y", params[1])
		parameters.Add("type", params[2])
		parameters.Add("page", params[3])
	case "id":
		parameters.Add("i", params[0])
		parameters.Add("plot", plot)
		parameters.Add("tomatoes", tomatoes)
	}
	URL.RawQuery = parameters.Encode()
	res, err := http.Get(URL.String())
	if err != nil {
		return nil, err
	}

	return res, nil
}
