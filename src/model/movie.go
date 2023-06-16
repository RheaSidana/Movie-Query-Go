package model

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Tconst         string `gorm:"primaryKey" json:"tconst" binding:"required"`
	TitleType      string `json:"titleType" binding:"required"`
	PrimaryTitle   string `json:"primaryTitle" binding:"required"`
	RuntimeMinutes int    `json:"runtineMinutes" binding:"required"`
	Genres         string `json:"genres" binding:"required"`
}
