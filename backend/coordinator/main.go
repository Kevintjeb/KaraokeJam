package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"gitlab.mi.hdm-stuttgart.de/mwa/ws18-karaoke/backend/common/rabbitmq"
	"gitlab.mi.hdm-stuttgart.de/mwa/ws18-karaoke/backend/coordinator/connection"
)

var addr = flag.String("addr", ":8080", "http service address")

func addCorsHeader(res http.ResponseWriter) {
	headers := res.Header()
	headers.Add("Access-Control-Allow-Origin", "*")
	headers.Add("Vary", "Origin")
	headers.Add("Vary", "Access-Control-Request-Method")
	headers.Add("Vary", "Access-Control-Request-Headers")
	headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
	headers.Add("Access-Control-Allow-Methods", "GET, POST,OPTIONS")
}
func main() {
	fmt.Println("Starting coordinator!")
	hub := connection.NewHub()

	router := mux.NewRouter().StrictSlash(true)

	router.Methods("OPTIONS").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			addCorsHeader(w)
			w.WriteHeader(http.StatusOK)
		})

	http.HandleFunc("/coordinator/ws", func(w http.ResponseWriter, r *http.Request) {
		connection.ServeWs(hub, w, r)
	})
	fmt.Println("Starting hub!")
	go hub.Run()

	conn := rabbitmq.Init()
	defer conn.Close()
	consumer, err := rabbitmq.NewConsumer(conn)

	if err != nil {
		fmt.Println(err)
	}

	messages := make(chan string)
	go consumer.Listen([]string{"queue"}, messages)

	go func() {
		for msg := range messages {
			var bytes = []byte(msg)
			var decodedSongMessage connection.SongMessage
			json.Unmarshal(bytes, &decodedSongMessage)
			hub.Message <- decodedSongMessage
		}
	}()

	if err != nil {
		panic(err)
	}

	fmt.Println("Listening on :8080")
	error := http.ListenAndServe(*addr, nil)
	if error != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
