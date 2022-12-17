package main

import (
	"github.com/Eiliv17/GinLibraryAPI/controllers"
	"github.com/Eiliv17/GinLibraryAPI/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	v1 := r.Group("/books")
	{
		// get the full list of books
		v1.GET("", controllers.GetBooks)

		// get a single book by ID
		v1.GET("/:id", controllers.GetBook)

		// create a book
		v1.POST("", controllers.CreateBook)
	}

	r.Run() // listen and serve on localhost with port defined in .env
}
