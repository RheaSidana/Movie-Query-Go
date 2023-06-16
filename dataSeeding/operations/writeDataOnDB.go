package operations

import (
	"Movies/initializer"
	"Movies/src/model"
	"fmt"
)

func WriteMovies(movies []model.Movie) {
	for i, movie := range movies {
		var dbMovie model.Rating
		if initializer.Db.Where("tconst=?", movie.Tconst).Find(&dbMovie).RowsAffected == 1 {
			fmt.Println("Movie found : ", i)
			continue
		}
		initializer.Db.Create(&movie)
	}

	fmt.Println("............Write Movies Successfully............")
}

func WriteRatings(ratings []model.Rating) {
	for i, rating := range ratings {
		var dbRating model.Rating
		if initializer.Db.Where("tconst=?", rating.Tconst).Find(&dbRating).RowsAffected == 1 {
			fmt.Println("Rating found : ", i)
			continue
		}
		initializer.Db.Create(&rating)
	}

	fmt.Println("............Write Ratings Successfully............")
}
