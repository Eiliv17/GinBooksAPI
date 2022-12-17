package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// book struct for database
type Book struct {
	Title           string             `bson:"title" json:"title" binding:"required"`
	ObjectID        primitive.ObjectID `bson:"_id" json:"ID"`
	Authors         []string           `bson:"authors" json:"authors" binding:"required"`
	PublicationDate time.Time          `bosn:"publicationDate" json:"publicationDate" binding:"required"`
	Publisher       string             `bson:"publisher" json:"publisher" binding:"required"`
	Language        string             `bson:"language" json:"language" binding:"required"`
	ISBN13          string             `bson:"isbn13" json:"isbn13" binding:"required"`
}
