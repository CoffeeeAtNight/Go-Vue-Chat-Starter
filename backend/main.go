package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type msg struct {
	Username  string
	Message   string
	Timestamp int
}

var logger log.Logger = *log.Default()
var upgrader websocket.Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var activeUsers []*websocket.Conn

func wsHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(w, r, nil)
	checkErr(err)

	log.Println("Client connected successfully!")

	activeUsers = append(activeUsers, conn)

	reader(conn)
}

func reader(conn *websocket.Conn) {
	m := msg{}

	logger.Println("Active connections: ", len(activeUsers))

	defer func() {
		for i, u := range activeUsers {
			if u == conn {
				activeUsers = append(activeUsers[:i], activeUsers[i+1:]...)
				break
			}
		}

		conn.Close()

		log.Println("A client left the chat!")
		logger.Println("Active connections: ", len(activeUsers))
	}()

	for {
		err := conn.ReadJSON(&m)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("Unexpected close error: %v", err)
			}
			return
		}

		fmt.Printf("• From : %s\n• Message: %s\n----------------\n", m.Username, m.Message)

		for _, u := range activeUsers {
			if err := u.WriteJSON(m); err != nil {
				log.Println(err)
				return
			}
		}

	}
}

func checkErr(err error) {
	if err != nil {
		logger.Println(err)
	}
}

func main() {
	logger.Println("Starting Webserver...")

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", wsHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
