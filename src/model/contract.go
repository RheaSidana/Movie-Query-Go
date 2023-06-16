package model

type ErrorResponse struct {
	Message string
}

type SuccessResponse struct {
	Message string
}

type LongDurMovies struct {
	Tconst         string `json:"tconst" binding:"required"`
	PrimaryTitle   string `json:"primaryTitle" binding:"required"`
	RuntimeMinutes int    `json:"runtineMinutes" binding:"required"`
	Genres         string `json:"genres" binding:"required"`
}

type LongDurMoviesTop10 struct {
	Movies []LongDurMovies
}

type TopRatedMovies struct {
	Tconst        string  `json:"tconst" binding:"required"`
	PrimaryTitle  string  `json:"primaryTitle" binding:"required"`
	Genres        string  `json:"genres" binding:"required"`
	AverageRating float32 `json:"averageRating" binding:"required"`
}

type TopRatedMoviesGT6 struct {
	Movies []TopRatedMovies
}

type MovieWithNumVotesGenres struct {
	Genres       string `json:"genres" binding:"required"`
	PrimaryTitle string `json:"primaryTitle" binding:"required"`
	NumVotes     int64  `json:"numVotes" binding:"required"`
}

type MovieWithNumVotes struct {
	PrimaryTitle string `json:"primaryTitle" binding:"required"`
	NumVotes     int64  `json:"numVotes" binding:"required"`
}

type SubTotalGenre struct {
	Genres string `json:"genres" binding:"required"`
	Total  int64  `json:"total" binding:"required"`
}

type SubTotalGenreMovie struct {
	Genres string `json:"genres" binding:"required"`
	Movies []MovieWithNumVotes
	Total  int64 `json:"total" binding:"required"`
}

type SubTotalGenreMovieList struct {
	MovieList []SubTotalGenreMovie
}

type Genres struct {
	Action      int `json:"Action"`
	Animation   int `json:"Animation"`
	Comedy      int `json:"Comedy"`
	Documentary int `json:"Documentary"`
	Drama       int `json:"Drama"`
	Family      int `json:"Family"`
	Fantasy     int `json:"Fantasy"`
	Horror      int `json:"Horror"`
	News        int `json:"News"`
	Romance     int `json:"Romance"`
	Short       int `json:"Short"`
	Sport       int `json:"Sport"`
	Rest        int `json:"Rest"`
}
