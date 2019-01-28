package main

import (
	"testing"

	"github.com/go-redis/redis"
)

func TestRoomService_Create_Leave(t *testing.T) {

	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	roomService = RoomService{Redis: RedisImpl{
		c: redisClient,
	}}

	t.Run("Create room, leave room", func(t *testing.T) {
		response := roomService.Create()
		_, err := roomService.Leave(response.User)

		if err != nil {
			t.FailNow()
		}

		resp, err := redisClient.SMembers(response.Room.Key).Result()

		if err != nil {
			t.FailNow()
		}

		if len(resp) != 0 {
			t.FailNow()
		}
	})
}


func TestRoomService_Join_leave(t *testing.T) {

	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	roomService = RoomService{Redis: RedisImpl{
		c: redisClient,
	}}

	t.Run("Create room, Join room, leave room", func(t *testing.T) {
		response := roomService.Create()
		var err error
		user, err := roomService.Join(Room{response.Room.Key})
		_, err = roomService.Join(Room{response.Room.Key})
		_, err = roomService.Join(Room{response.Room.Key})
		leftUser, err := roomService.Leave(response.User)

		if err != nil {
			t.FailNow()
		}

		if leftUser.Success != true {
			t.FailNow()
		}

		if user.Room.Key == "" || user.User.Key == "" {
			t.FailNow()
		}

		resp, err := redisClient.SMembers(response.Room.Key).Result()

		if err != nil {
			t.FailNow()
		}

		if len(resp) != 3 {
			t.FailNow()
		}
	})
}

func TestRoomService_Create(t *testing.T) {

	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	roomService = RoomService{Redis: RedisImpl{
		c: redisClient,
	}}

	t.Run("Create room", func(t *testing.T) {
		response := roomService.Create()

		resp, err := redisClient.SMembers(response.Room.Key).Result()

		if err != nil {
			t.FailNow()
		}

		if response.Room.Key == "" || response.User.Key == "" {
			t.FailNow()
		}

		if len(resp) != 1 {
			t.FailNow()
		}
	})
}
