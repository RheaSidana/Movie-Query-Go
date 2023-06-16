package model

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	Tconst        string  `gorm:"references:Movies.tconst" json:"tconst" binding:"required"`
	AverageRating float32 `json:"averageRating" binding:"required"`
	NumVotes      int64   `json:"numVotes" binding:"required"`
}
