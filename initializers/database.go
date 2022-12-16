package initializers

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Client

func ConnectToDB() {

	// Getting the Database URI from the environment variables
	URI := os.Getenv("MONGODB_URI")
	if URI == "" {
		log.Fatal("You must set your 'DB_URI' environment variable")
	}

	// Connecting to the database
	DB, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal("Error connecting to Database")
	}
	/* defer func() {
		if err = DB.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}() */

	// Pings the database as a test
	if err := DB.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
}
