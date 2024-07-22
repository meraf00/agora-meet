package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Message struct {
	UserId  string `json:"userId"`
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var meetings = make(map[string][]*websocket.Conn)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	params := mux.Vars(r)
	meetingId, ok := params["meetingId"]

	if !ok {
		fmt.Println("Invalid meeting id")
		return
	}

	clients[conn] = true
	meetings[meetingId] = append(meetings[meetingId], conn)

	fmt.Println(clients)
	fmt.Println(meetings)

	for {
		var msg Message

		err := conn.ReadJSON(&msg)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(msg.Type, msg.UserId)

		for _, c := range meetings[meetingId] {
			if c != conn && clients[c] {
				err := c.WriteJSON(msg)

				if err != nil {
					fmt.Println("Error:", err)
					delete(clients, c)
				}
			}
		}

	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/ws/{meetingId}", handleConnections)

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}
