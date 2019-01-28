package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"gitlab.mi.hdm-stuttgart.de/mwa/ws18-karaoke/backend/common/rabbitmq"
)

var redisClient *redis.Client
var roomService RoomService

type joinRoomRequest struct {
	RoomKey string
}

type leaveRoomRequest struct {
	UserKey string
}

type RedisClient interface {
	SRem(key string, members string)
	Set(key string, value string, expiration time.Duration)
	SAdd(key string, members string)
	Expire(key string, expiration time.Duration)
	Del(keys string)
	Get(key string) (string, error)
	Exists(keys string) (int64, error)
}

type rabbitMqMessage struct {
	Type    string
	Payload interface{}
}

var emitter rabbitmq.Emitter

func main() {
	conn := rabbitmq.Init()
	defer conn.Close()

	var err error
	emitter, err = rabbitmq.NewEmitter(conn)
	if err != nil {
		panic(err)
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	roomService = RoomService{Redis: RedisImpl{
		c: redisClient,
	}}

	router := mux.NewRouter().StrictSlash(true)

	router.Methods("OPTIONS").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			addCorsHeader(w)
			w.WriteHeader(http.StatusOK)
		})

	router.HandleFunc("/rooms/new", createRoom)
	router.HandleFunc("/rooms/join", joinRoom).Methods("POST")
	router.HandleFunc("/rooms/leave", leaveRoom).Methods("POST")

	fmt.Println("Listening on :8081")

	log.Fatal(http.ListenAndServe(":8081", router))

}

func addCorsHeader(res http.ResponseWriter) {
	headers := res.Header()
	headers.Add("Access-Control-Allow-Origin", "*")
	headers.Add("Vary", "Origin")
	headers.Add("Vary", "Access-Control-Request-Method")
	headers.Add("Vary", "Access-Control-Request-Headers")
	headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
	headers.Add("Access-Control-Allow-Methods", "GET, POST,OPTIONS")
}

func createRoom(writer http.ResponseWriter, request *http.Request) {
	room := roomService.Create()
	message := rabbitMqMessage{
		Type:    "NEW_ROOM",
		Payload: room.Room.Key,
	}

	marshalled, _ := json.Marshal(message)
	_ = emitter.Push(string(marshalled), "rooms")

	log.Print(marshalled)

	writer.Header().Set("Content-Type", "application/json")
	addCorsHeader(writer)
	writer.WriteHeader(http.StatusCreated)
	e := writeJSON(writer, room)

	if e != nil {
		fmt.Print(fmt.Errorf("error creating new Room %s", e))
	} else {
		log.Println("Successfully send data!")
	}
}

func joinRoom(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var joinRequest joinRoomRequest

	if err := decoder.Decode(&joinRequest); err != nil {
		fmt.Print(fmt.Errorf("error parsing body for joinRoom %s", err))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := roomService.Join(Room{Key: joinRequest.RoomKey})
	addCorsHeader(writer)
	if err != nil {
		log.Print(fmt.Errorf("error joining room  %s", err))
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.WriteHeader(http.StatusOK)
	}

	e := writeJSON(writer, response)

	if e != nil {
		log.Print(fmt.Errorf("error sending data %s", e))
	} else {
		log.Println("Successfully send data!")
	}
}

func leaveRoom(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var leaveRoomRequest leaveRoomRequest

	if err := decoder.Decode(&leaveRoomRequest); err != nil {
		log.Print(fmt.Errorf("error parsing body for leaveRoom %s", err))
	}

	leaveResponse, err := roomService.Leave(User{Key: leaveRoomRequest.UserKey})
	addCorsHeader(writer)
	if err != nil {
		log.Print(fmt.Errorf("error leaving room  %s", err))
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.WriteHeader(http.StatusOK)
	}
	e := writeJSON(writer, leaveResponse)

	if e != nil {
		log.Println(fmt.Errorf("error sending data %s", e))
	} else {
		log.Println("Successfully send data!")
	}
}

func writeJSON(writer http.ResponseWriter, leaveResponse interface{}) error {
	e := json.NewEncoder(writer).Encode(leaveResponse)
	return e
}
