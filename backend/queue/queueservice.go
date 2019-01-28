package main

import (
	"context"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type Track struct {
	ID         string   `json:"id,omitempty" bson:"id"`
	Name       string   `json:"name,omitempty" bson:"name"`
	Artist     []Artist `json:"artist,omitempty" bson:"artist"`
	Album      Album    `json:"album,omitempty" bson:"album"`
	Popularity int      `json:"popularity,omitempty" bson:"popularity"`
	Duration   int      `json:"duration,omitempty" bson:"duration"`
}

type Artist struct {
	ID    string `json:"id,omitempty" bson:"id"`
	Name  string `json:"name,omitempty" bson:"name"`
	Image Image  `json:"image,omitempty" bson:"image"`
}
type Album struct {
	ID     string   `json:"id,omitempty" bson:"id"`
	Name   string   `json:"name,omitempty" bson:"name"`
	Artist []Artist `json:"artist,omitempty" bson:"artist"`
	Cover  Image    `json:"cover,omitempty" bson:"cover"`
}

type Image struct {
	Small  string `json:"64,omitempty" bson:"small"`
	Medium string `json:"300,omitempty" bson:"medium"`
	Large  string `json:"640,omitempty" bson:"large"`
}

type Queue struct {
	RoomKey string  `bson:"RoomKey"`
	Tracks  []Track `bson:"Tracks"`
}

type QueueService interface {
	Add(track Track, roomKey string) error
	Remove(trackId string, roomKey string) error
	Get(roomKey string) ([]Track, error)
	CreateRoom(roomKey interface{}) error
}

type QueueServiceImpl struct {
	collection *mongo.Collection
}

func (service QueueServiceImpl) CreateRoom(roomkey interface{}) error {
	roomkey = roomkey.(string)
	log.Printf("Creating queue storage room for %s", roomkey)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, e := service.collection.InsertOne(ctx, bson.D{{"RoomKey", roomkey}, {"Tracks", bson.A{}}})

	if e != nil {
		log.Printf("Error creating a room in mongodb %s", e)
	}

	return e
}

func (service QueueServiceImpl) Add(track Track, roomKey string) error {

	var room Queue
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	filterDoc := bson.D{{"RoomKey", bson.D{{
		"$eq",
		roomKey,
	}}}}
	err := service.collection.FindOne(ctx, filterDoc).Decode(&room)

	if err != nil {
		log.Printf("Error while finding a document %s", err)
		return err
	}
	_, err = service.collection.UpdateOne(ctx, filterDoc, bson.D{
		{"$push",
			bson.D{{"Tracks", track}}}},
	)

	if err != nil {
		log.Printf("Error while updating a document %s", err)
		return err
	}

	return nil
}

func (service QueueServiceImpl) Remove(trackId string, roomKey string) error {
	var room Queue
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	filterDoc := bson.D{{"RoomKey", bson.D{{
		"$eq",
		roomKey,
	}}}}
	err := service.collection.FindOne(ctx, filterDoc).Decode(&room)

	if err != nil {
		log.Printf("Error while finding a document %s", err)
		return err
	}

	_, err = service.collection.UpdateOne(ctx, filterDoc, bson.D{
		{"$pull",
			bson.D{{"Tracks", bson.D{
				{"id", trackId}},
			}},
		}},
	)

	if err != nil {
		log.Printf("Error while removing a document %s", err)
		return err
	}

	return nil
}

func (service QueueServiceImpl) Get(roomKey string) ([]Track, error) {
	var room Queue

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	filter := bson.D{{"RoomKey", roomKey}}

	//err := collection.FindOne(ctx, filter).Decode(&room)
	err := service.collection.FindOne(ctx, filter).Decode(&room)

	if err != nil {
		log.Printf("Error while finding a document %s", err)
		return nil, err
	}

	return room.Tracks, nil
}
