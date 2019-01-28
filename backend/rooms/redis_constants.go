package main

import (
	"time"

	"github.com/satori/go.uuid"
)

const ROOM_ADMIN_POSTFIX = "ROOM_ADMIN:"
const ROOM_KEY_PREFIX = "ROOM_KEY:"
const USER_KEY_PREFIX = "USER_KEY:"
const DEFAULT_KEY_EXPIRATION = time.Hour * 12

func generateRoomAdminKey(roomkey string) string {
	return ROOM_ADMIN_POSTFIX + roomkey
}

func generateRoomKey() string {
	roomKey := uuid.NewV4()

	return ROOM_KEY_PREFIX + roomKey.String()
}

func generateUserKey() string {
	userKey := uuid.NewV4()

	return USER_KEY_PREFIX + userKey.String()
}
