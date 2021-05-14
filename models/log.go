package models

import (
	"context"
	"time"

	"github.com/michaelt0520/nfc-card/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SQLLog struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserID    uint               `bson:"user_id"`
	Action    string             `bson:"action"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func CreateIssue(log SQLLog) error {
	//Get MongoDB connection using connectionhelper.
	client := db.GetMongoClient()

	//Create a handle to the respective collection in the database.
	collection := client.Database(db.DB).Collection(db.ISSUES)

	//Perform InsertOne operation & validate against the error.
	if _, err := collection.InsertOne(context.TODO(), log); err != nil {
		return err
	}

	//Return success without any error.
	return nil
}
