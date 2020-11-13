package dbHandler

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Item that MongoDB's document should save
type Item struct {
	UserName   string
	DeviceID   string
	Brightness string
	CCT        string
	Timer      string
	UsedTime   string
}

// Storage of Mongo interface
type Storage interface {
	GetByName(context.Context, string) (*Item, error)
	Put(context.Context, *Item) error
}

// MongoStorage : MongoDB configuration
type MongoStorage struct {
	*mongo.Collection
	DB string
}

func dbInit() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())

}

func insertDoc(client *Client) {
	collection := client.Database("testing").Collection("numbers")
}

func dbInsert(collection *Collection) {
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
}
