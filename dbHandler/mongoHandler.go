package dbHandler

import (
	"context"
	"time"

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
	PurgeByName(context.Context, string) error
}

// Mongodb : MongoDB configuration
type Mongodb struct {
	Client     *mongo.Client
	DB         string
	Collection string
}

func dbInit(_db string, _col string) (*Mongodb, error) {
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

	db := Mongodb{
		Client:     client,
		DB:         _db,
		Collection: _col,
	}

	return &db, nil
}
