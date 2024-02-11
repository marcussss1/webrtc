package main

import (
	"log"
	"net/http"
)

func main() {
	AllRooms.Init()

	http.HandleFunc("/create", CreateRoomRequestHandler)
	http.HandleFunc("/join", JoinRoomRequestHandler)

	log.Println("Starting Server on Port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
