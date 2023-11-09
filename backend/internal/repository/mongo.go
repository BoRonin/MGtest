package repository

import (
	"context"
	"fmt"
	"log"
	"mgtest/internal/models"
	"mgtest/internal/storage/mongo"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mongo struct {
	C *mongo.Mongo
}

func NewMongo(client *mongo.Mongo) *Mongo {
	return &Mongo{
		C: client,
	}
}

func (m *Mongo) InsertProfile(ctx context.Context, data models.Data) (models.Data, error) {
	log.Println("inside InsertProfile")
	coll := m.C.Client.Database("Data").Collection("Profiles")
	result, err := coll.InsertOne(ctx, data)
	if err != nil {
		return models.Data{}, err
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		data.ID = oid.Hex()
	}
	return data, nil
}

func (m *Mongo) UpdateProfile(ctx context.Context, data models.Data, prof models.Data) (int, error) {
	log.Println("inside UpdateProfile")
	coll := m.C.Client.Database("Data").Collection("Profiles")
	log.Println("data id:", data.ID)
	id, _ := primitive.ObjectIDFromHex(data.ID)
	log.Printf("new id:%v\n", id)
	fields, delta := getDiff(data, prof)
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", fields}}
	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}

	return delta, nil
}
func (m *Mongo) GetProfile(ctx context.Context, idString string) (models.Data, error) {
	log.Println("inside GetProfile")
	coll := m.C.Client.Database("Data").Collection("Profiles")
	var data models.Data
	id, _ := primitive.ObjectIDFromHex(idString)
	filter := bson.D{{"_id", id}}
	if err := coll.FindOne(ctx, filter).Decode(&data); err != nil {
		return models.Data{}, err
	}
	return data, nil
}

func (m *Mongo) Close() {
	err := m.C.Client.Disconnect(context.Background())
	if err != nil {
		log.Printf("couldn't disconnect from db: %v", err)
	}
}

func getDiff(new models.Data, old models.Data) ([]bson.E, int) {
	var fields []bson.E
	count := 0
	if new.Name != old.Name {
		fields = append(fields, bson.E{Key: "name", Value: new.Name})
		count++
	}
	if !reflect.DeepEqual(new.Bags, old.Bags) {
		for i, b := range new.Bags {
			for j, f := range b.Facts {
				if f.Info != old.Bags[i].Facts[j].Info {
					fields = append(fields, bson.E{Key: fmt.Sprintf("bags.%d.facts.%d.info", i, j), Value: f.Info})
					count++
				}
				if f.InterestingNumber != old.Bags[i].Facts[j].InterestingNumber {
					fields = append(fields, bson.E{Key: fmt.Sprintf("bags.%d.facts.%d.number", i, j), Value: f.InterestingNumber})
					count++
				}
			}
		}
	}
	return fields, count
}
