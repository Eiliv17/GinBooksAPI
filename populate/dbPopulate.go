package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/Eiliv17/GinBooksAPI/initializers"
	"github.com/Eiliv17/GinBooksAPI/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// struct to decode json
	booksData := []struct {
		Title           string   `json:"title" binding:"required"`
		Authors         []string `json:"authors" binding:"required"`
		PublicationDate string   `json:"publicationDate" binding:"required"`
		Publisher       string   `json:"publisher" binding:"required"`
		Language        string   `json:"language" binding:"required"`
		ISBN13          string   `json:"isbn13" binding:"required"`
	}{}

	data, err := os.ReadFile("./populate/dbData.json")
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(data, &booksData)

	// gets the database name and collection name from env variables
	dbname := os.Getenv("DB_NAME")
	collname := os.Getenv("COLL_NAME")
	coll := initializers.DB.Database(dbname).Collection(collname)

	// inserts the books from json
	for _, book := range booksData {
		pubdate, err := time.Parse("2006-01-02", book.PublicationDate)
		if err != nil {
			log.Fatal(err)
		}

		_, err = coll.InsertOne(context.TODO(), models.Book{
			Title:           book.Title,
			Authors:         book.Authors,
			PublicationDate: pubdate,
			Publisher:       book.Publisher,
			Language:        book.Language,
			ISBN13:          book.ISBN13,
			ObjectID:        primitive.NewObjectIDFromTimestamp(time.Now()),
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}
