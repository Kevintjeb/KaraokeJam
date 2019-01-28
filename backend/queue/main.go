package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mongodb/mongo-go-driver/mongo"
	"gitlab.mi.hdm-stuttgart.de/mwa/ws18-karaoke/backend/common/rabbitmq"
)

type addSongRequest struct {
	Song    Track
	RoomKey string
	UserKey string
}

type removeSongRequest struct {
	Song    Track
	RoomKey string
}

type rabbitMqMessage struct {
	Type    string
	Payload string
}

type SongMessage struct {
	Type    string
	RoomKey string
	Track   Track
}

var service QueueService
var emitter rabbitmq.Emitter
var consumer rabbitmq.Consumer

func main() {
	conn := rabbitmq.Init()
	defer conn.Close()

	var err error

	emitter, err = rabbitmq.NewEmitter(conn)
	consumer, err = rabbitmq.NewConsumer(conn)

	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, "mongodb://mongodb:27017")

	if err != nil {
		panic(err)
	}

	collection := client.Database("queue").Collection("rooms")

	service = QueueServiceImpl{
		collection: collection,
	}

	router := mux.NewRouter().StrictSlash(true)

	router.Methods("OPTIONS").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			addCorsHeader(w)
			w.WriteHeader(http.StatusOK)
		})

	router.HandleFunc("/songs", addSong).Methods("POST")
	router.HandleFunc("/songs/delete", removeSong).Methods("POST")
	router.HandleFunc("/songs/{RoomKey}", getAllSongs).Methods("GET")

	log.Println("Listening on :8082")

	messages := make(chan string)
	go consumer.Listen([]string{"rooms"}, messages)

	go func() {
		for msg := range messages {
			var message rabbitMqMessage
			err := json.Unmarshal([]byte(msg), &message)

			if err != nil {
				log.Printf("Error! %s", err)
			} else {
				_ = service.CreateRoom(message.Payload)
			}
		}
	}()

	log.Fatal(http.ListenAndServe(":8082", router))
}

func addCorsHeader(res http.ResponseWriter) {
	headers := res.Header()
	headers.Add("Access-Control-Allow-Origin", "*")
	headers.Add("Vary", "Origin")
	headers.Add("Vary", "Access-Control-Request-Method")
	headers.Add("Vary", "Access-Control-Request-Headers")
	headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
	headers.Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")
}

func getAllSongs(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	tracks, e := service.Get(params["RoomKey"])
	addCorsHeader(writer)
	if e != nil {
		log.Printf("Error retrieving all songs %s", e)
		writer.WriteHeader(http.StatusInternalServerError)
	}

	err := json.NewEncoder(writer).Encode(tracks)

	if err != nil {
		log.Print(fmt.Errorf("error sending data %s", err))
	} else {
		log.Println("Successfully send data!")
	}
}

func removeSong(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var removeSong removeSongRequest
	addCorsHeader(writer)

	if err := decoder.Decode(&removeSong); err != nil {
		log.Println(fmt.Errorf("error parsing body for removeSong %s", err))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := service.Remove(removeSong.Song.ID, removeSong.RoomKey)

	if err != nil {
		log.Println(fmt.Errorf("error removing song %s", err))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, e := json.Marshal(SongMessage{
		Type:    "REMOVE_SONG",
		RoomKey: removeSong.RoomKey,
		Track:   removeSong.Song,
	})

	if e != nil {
		log.Println(fmt.Errorf("error marshalling removeSong %s", e))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	_ = emitter.Push(string(bytes), "queue")
}

func addSong(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var songRequest addSongRequest
	addCorsHeader(writer)
	if err := decoder.Decode(&songRequest); err != nil {
		fmt.Print(fmt.Errorf("error parsing body for songRequest %s", err))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := service.Add(songRequest.Song, songRequest.RoomKey)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, e := json.Marshal(SongMessage{
		Type:    "ADD_SONG",
		RoomKey: songRequest.RoomKey,
		Track:   songRequest.Song,
	})

	if e != nil {
		fmt.Print(fmt.Errorf("error marshalling songrequest %s", e))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	_ = emitter.Push(string(bytes), "queue")
}
