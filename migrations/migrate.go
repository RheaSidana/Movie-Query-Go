package main

import (
	"Movies/initializer"
	"Movies/src/model"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {
	initializer.Db.AutoMigrate(&model.Movie{})
	initializer.Db.AutoMigrate(&model.Rating{})
}
