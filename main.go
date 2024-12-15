package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type webSocketHandler struct {
	upgrader websocket.Upgrader
}

func(wsHandler webSocketHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	connection, err := wsHandler.upgrader.Upgrade(res, req, nil)

	if err != nil {
		log.Printf("error %s when upgrading connection to websocket", err)

		return
	}

	defer connection.Close()
}

func main() {
	webSocketHandler := webSocketHandler{
		upgrader: websocket.Upgrader{},
	}

	http.Handle("/", webSocketHandler)
	log.Print("Starting server...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}