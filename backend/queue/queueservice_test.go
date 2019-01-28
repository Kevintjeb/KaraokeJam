package main

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"testing"
	"time"
)

func TestQueueService_Create(t *testing.T) {
	var roomKey = "test-key-1"
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, "mongodb://localhost:27017")

	if err != nil {
		panic(err)
	}

	collection := client.Database("queue").Collection("rooms")

	service = QueueServiceImpl{
		collection: collection,
	}

	t.Run("Create queue in mongodb", func(t *testing.T) {
		err := service.CreateRoom(roomKey)

		if err != nil {
			t.FailNow()
		}

		result, err := service.Get(roomKey)

		if len(result) != 0{
			t.FailNow()
		}

	})
}

func TestQueueService_Add_Songs(t *testing.T) {
	var roomKey = "test-key-2"
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, "mongodb://localhost:27017")

	if err != nil {
		panic(err)
	}

	collection := client.Database("queue").Collection("rooms")

	service = QueueServiceImpl{
		collection: collection,
	}

	t.Run("Create queue in mongodb, add 2 songs", func(t *testing.T) {
		err := service.CreateRoom(roomKey)

		if err != nil {
			t.FailNow()
		}

		_ = service.Add(Track{}, roomKey)
		_ = service.Add(Track{}, roomKey)

		result, err := service.Get(roomKey)

		if len(result) != 2 {
			t.FailNow()
		}

	})
}


func TestQueueService_Delete_Song(t *testing.T) {
	var roomKey = "test-key-3"
	var deleteKey = "DELETE-ME"
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, "mongodb://localhost:27017")

	if err != nil {
		panic(err)
	}

	collection := client.Database("queue").Collection("rooms")

	service = QueueServiceImpl{
		collection: collection,
	}

	t.Run("Create queue in mongodb, add 2 songs, delete 1 song", func(t *testing.T) {
		err := service.CreateRoom(roomKey)

		if err != nil {
			t.FailNow()
		}

		_ = service.Add(Track{}, roomKey)
		_  = service.Add(Track{
			ID: deleteKey,
		}, roomKey)

		result, err := service.Get(roomKey)

		if len(result) != 2 {
			t.FailNow()
		}

		err = service.Remove(deleteKey,roomKey)

		if err != nil {
			t.FailNow()
		}
		result, err = service.Get(roomKey)

		if err != nil {
			t.FailNow()
		}
		if len(result) != 1 {
			t.FailNow()
		}
	})
}
