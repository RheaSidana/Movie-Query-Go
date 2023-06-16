package main

import (
	op "Movies/dataSeeding/operations"
	"Movies/initializer"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {
	records := op.ReadCSV(`dataSeeding/data/movies.csv`)
	movies := op.CreateMovies(records)

	records = op.ReadCSV(`dataSeeding/data/ratings.csv`)
	ratings := op.CreateRatings(records)

	op.WriteMovies(movies)
	op.WriteRatings(ratings)
}
