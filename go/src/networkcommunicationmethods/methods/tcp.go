package methods

import (
	"fmt"
	"log"
	"net"

	"sync/atomic"
)

var count int64

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024) // 고정 크기 버퍼
	n, err := conn.Read(buffer)
	if err != nil {
		log.Println("Error reading from TCP: ", err)
		return
	}
	fmt.Println("Server: Received: ", string(buffer[:n]))

	atomic.StoreInt64(&count, 1)
	message := fmt.Sprintf("Server: msg received, %d", count)
	conn.Write([]byte(message))
}

func RunTCP() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	go func() {

		// server
		go func() {
			for {
				conn, err := listener.Accept()
				if err != nil {
					log.Fatal(err)
				}

				go handleTCPConnection(conn)
			}
		}()

		// client
		go func() {
			conn, err := net.Dial("tcp", ":8080")
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()

			message := "Hello\nTCP!\n"
			_, err = conn.Write([]byte(message))
			if err != nil {
				log.Fatal(err)
			}

			buffer := make([]byte, 1024) // 고정 크기 버퍼
			n, _ := conn.Read(buffer)
			fmt.Println("Client: Server response: ", string(buffer[:n]))
		}()
	}()
}
