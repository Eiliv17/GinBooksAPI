package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/Eiliv17/GinLibraryAPI/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

// custom time marshaler and unmarshaler
type CustomTime struct {
	time.Time
}

func (t CustomTime) MarshalJSON() ([]byte, error) {
	out := t.Time.Format("2006-01-02")
	return []byte(`"` + out + `"`), nil
}

func (t *CustomTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	nt, err := time.Parse(`"`+"2006-01-02"+`"`, string(b))
	if err != nil {
		return err
	}
	*t = CustomTime{nt}
	return nil
}

// book struct for database
type Book struct {
	Title           string     `bson:"title" json:"title"`
	Authors         []string   `bson:"authors" json:"authors"`
	PublicationDate CustomTime `bosn:"publicationDate" json:"publicationDate"`
	Publisher       string     `bson:"publisher" json:"publisher"`
	Language        string     `bson:"language" json:"language"`
	ISBN13          string     `bson:"isbn13" json:"isbn13"`
}

func main() {
	data, err := os.ReadFile("./populate/dbData.json")
	if err != nil {
		log.Fatal(err)
	}

	var books []Book
	json.Unmarshal(data, &books)

	// gets the database name and collection name from env variables
	dbname := os.Getenv("DB_NAME")
	collname := os.Getenv("COLL_NAME")
	coll := initializers.DB.Database(dbname).Collection(collname)

	// inserts the books from json
	for _, book := range books {
		_, err := coll.InsertOne(context.TODO(), book)
		if err != nil {
			log.Fatal(err)
		}
	}
}
