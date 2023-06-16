package operations

import (
	"Movies/src/model"
	"fmt"
	"strconv"
)

func mapToMovie(index int, value string, movie model.Movie) model.Movie {
	switch index {
	case 0:
		movie.Tconst = value
	case 1:
		movie.TitleType = value
	case 2:
		movie.PrimaryTitle = value
	case 3:
		movie.RuntimeMinutes, _ = strconv.Atoi(value)
	case 4:
		movie.Genres = value
	}

	return movie
}

func createMovie(record []string) model.Movie {

	var movie model.Movie

	for i, field := range record {
		movie = mapToMovie(i, string(field), movie)
	}

	return movie
}

func CreateMovies(records [][]string) []model.Movie {
	var movies []model.Movie

	for i, record := range records {
		if i == 0 {
			continue
		}
		movie := createMovie(record)
		movies = append(movies, movie)
	}

	fmt.Println("............Create Movies Successfully............")

	return movies
}
