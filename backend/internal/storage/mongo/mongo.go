package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client *mongo.Client
}

func New(url string, user string, pass string) *Mongo {
	clientOptions := options.Client().ApplyURI(url)
	clientOptions.SetAuth(options.Credential{
		Username: user,
		Password: pass,
	})
	c, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("could not connect to mongo: %s", err)
	}
	log.Println("Connected to mongo")
	return &Mongo{
		Client: c,
	}
}
