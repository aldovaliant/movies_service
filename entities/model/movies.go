package model

// QueryData is the type to create the search query
type QueryData struct {
	Title      string
	Year       string
	ImdbId     string
	SearchType string
	Page       string
}

//SearchResult is the type for the search results
type SearchResult struct {
	Title  string
	Year   string
	ImdbID string
	Type   string
}

//SearchResponse is the struct of the response in a search
type SearchResponse struct {
	Search       []SearchResult
	Response     string
	Error        string
	totalResults int
}

type MovieResult struct {
	Title             string
	Year              string
	Rated             string
	Released          string
	Runtime           string
	Genre             string
	Director          string
	Writer            string
	Actors            string
	Plot              string
	Language          string
	Country           string
	Awards            string
	Poster            string
	Metascore         string
	ImdbRating        string
	ImdbVotes         string
	ImdbID            string
	Type              string
	TomatoMeter       string
	TomatoImage       string
	TomatoRating      string
	TomatoReviews     string
	TomatoFresh       string
	TomatoRotten      string
	TomatoConsensus   string
	TomatoUserMeter   string
	TomatoUserRating  string
	TomatoUserReviews string
	TomatoURL         string
	DVD               string
	BoxOffice         string
	Production        string
	Website           string
	Response          string
	Error             string
}
