package controllers

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/Eiliv17/GinBooksAPI/initializers"
	"github.com/Eiliv17/GinBooksAPI/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// function to retrieve all books
func GetBooks(c *gin.Context) {
	// gets the database info
	dbname := os.Getenv("DB_NAME")
	collname := os.Getenv("COLL_NAME")
	coll := initializers.DB.Database(dbname).Collection(collname)

	// searches the database for the book
	result, _ := coll.Find(context.TODO(), bson.D{})

	// populates the slice
	var books []models.Book
	result.All(context.TODO(), &books)

	if len(books) == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{})
		return
	}

	c.IndentedJSON(http.StatusOK, books)
}

// function to retrieve a book from the database given an id
func GetBook(c *gin.Context) {
	// gets the id from the URL parameters
	bookID := c.Param("id")
	if bookID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty request",
		})
		return
	}

	objID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid book ID",
		})
		return
	}

	// gets the database info
	dbname := os.Getenv("DB_NAME")
	collname := os.Getenv("COLL_NAME")
	coll := initializers.DB.Database(dbname).Collection(collname)

	// searches the database for the book
	result := coll.FindOne(context.TODO(), bson.M{"_id": objID})

	var book models.Book
	err = result.Decode(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No book with that ID",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

// function that creates a book from the request body data
func CreateBook(c *gin.Context) {
	// body struct
	bodyData := struct {
		Title           string   `json:"title" binding:"required"`
		Authors         []string `json:"authors" binding:"required"`
		PublicationDate string   `json:"publicationDate" binding:"required"`
		Publisher       string   `json:"publisher" binding:"required"`
		Language        string   `json:"language" binding:"required"`
		ISBN13          string   `json:"isbn13" binding:"required"`
	}{}

	// get data from the request body
	if err := c.ShouldBindJSON(&bodyData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid form request",
		})
		return
	}
	pubdate, err := time.Parse("2006-01-02", bodyData.PublicationDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong date format use: YYYY-MM-DD",
		})
		return
	}

	// gets the database info
	dbname := os.Getenv("DB_NAME")
	collname := os.Getenv("COLL_NAME")
	coll := initializers.DB.Database(dbname).Collection(collname)

	// inserts the book into the database
	_, err = coll.InsertOne(context.TODO(), models.Book{
		Title:           bodyData.Title,
		Authors:         bodyData.Authors,
		PublicationDate: pubdate,
		Publisher:       bodyData.Publisher,
		Language:        bodyData.Language,
		ISBN13:          bodyData.ISBN13,
		ObjectID:        primitive.NewObjectIDFromTimestamp(time.Now()),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "error inserting the book",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"result": "book inserted successfully",
	})
}
