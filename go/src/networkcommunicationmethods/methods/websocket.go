package methods

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handler3(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Client connected")
	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		fmt.Println("Received:", string(msg))

		// 클라이언트에 응답 전송
		err = conn.WriteMessage(messageType, []byte("Hello from server!"))
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

func RunWebsocket() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", handler3)

	go func() {

		// server
		go func() {
			fmt.Println("Starting HTTP server on :8085")
			err := http.ListenAndServe(":8085", mux)
			if err != nil {
				log.Fatal(err)
			}
		}()

		// client
		go func() {
			// WebSocket 서버에 연결
			conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8085/ws", nil)
			if err != nil {
				log.Fatal("Failed to connect:", err)
			}
			defer conn.Close() // ✅ WebSocket 클라이언트는 Close() 필요

			// 서버에 메시지 전송
			message := "Hello, WebSocket Server!"
			err = conn.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Fatal("Write error:", err)
			}

			// 서버에서 응답 수신
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Fatal("Read error:", err)
			}

			fmt.Println("Server response:", string(msg))
		}()
	}()
}
