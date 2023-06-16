package main

import (
	"Movies/initializer"
	"Movies/src/modules/movies"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {
	r := gin.Default()
	
	movies.Apis(r)
	
	r.Run()
}

