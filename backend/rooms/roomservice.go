package main

import (
	"log"
)

type Room struct {
	Key string
}

type User struct {
	Key string
}

type CreateResponse struct {
	Room Room
	User User
}

type JoinResponse struct {
	Room Room
	User User
}

type LeaveResponse struct {
	Success bool
}

type RoomService struct {
	Redis RedisClient
}


func (service *RoomService) Create() CreateResponse {
	userKey := generateUserKey()
	roomKey := generateRoomKey()
	adminKey := generateRoomAdminKey(roomKey)

	service.Redis.Set(adminKey, userKey, DEFAULT_KEY_EXPIRATION)
	service.Redis.SAdd(roomKey, userKey)
	service.Redis.Expire(roomKey, DEFAULT_KEY_EXPIRATION)
	service.Redis.Set(userKey, roomKey, DEFAULT_KEY_EXPIRATION)

	createdRoom := CreateResponse{
		User: User{Key: userKey},
		Room: Room{Key: roomKey},
	}

	return createdRoom
}

func (service *RoomService) Leave(user User) (LeaveResponse, error) {
	RoomKey, err := service.Redis.Get(user.Key)

	if err != nil {
		log.Printf("Couldn't find a Room for User %s", user.Key)

		return LeaveResponse{Success: false}, nil
	}

	service.Redis.SRem(RoomKey, user.Key)
	service.Redis.Del(user.Key)

	return LeaveResponse{Success: true}, nil
}
func (service *RoomService) Join(room Room) (JoinResponse, error) {
	_, err := service.Redis.Exists(room.Key)

	if err != nil {
		return JoinResponse{}, err
	}

	userKey := generateUserKey()
	service.Redis.Set(userKey, room.Key, DEFAULT_KEY_EXPIRATION)
	service.Redis.SAdd(room.Key, userKey)

	response := JoinResponse{
		User: User{Key: userKey},
		Room: room,
	}

	return response, nil

}
