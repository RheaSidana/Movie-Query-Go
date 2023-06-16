package operations

import (
	"Movies/src/model"
	"fmt"
	"strconv"
)

func mapToRating(index int, value string, rating model.Rating) model.Rating {
	switch index {
	case 0:
		rating.Tconst = value
	case 1:
		val, _ := strconv.ParseFloat(value, 32)
		rating.AverageRating = float32(val)
	case 2:
		rating.NumVotes, _ = strconv.ParseInt(value, 10, 64)
	}

	return rating
}

func createRating(record []string) model.Rating {

	var rating model.Rating
	for i, field := range record {
		rating = mapToRating(i, string(field), rating)
	}

	return rating
}

func CreateRatings(records [][]string) []model.Rating {
	var ratings []model.Rating

	for i, record := range records {
		if i == 0 {
			continue
		}
		rating := createRating(record)
		ratings = append(ratings, rating)
	}

	fmt.Println("............Create Ratings Successfully............")

	return ratings
}
