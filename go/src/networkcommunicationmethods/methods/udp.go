package methods

import (
	"fmt"
	"log"
	"net"
	"sync/atomic"
)

var count2 int64

func handleUDPConnection(conn *net.UDPConn) {
	buffer := make([]byte, 1024) // 고정 크기 버퍼
	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println("Error reading from UDP: ", err)
			// time.Sleep(2 * time.Second)
			continue
		}
		fmt.Printf("Server: Received from %s: %s", clientAddr, string(buffer[:n]))

		atomic.StoreInt64(&count2, 1)
		message := fmt.Sprintf("Server: msg received, %d", count2)
		conn.WriteToUDP([]byte(message), clientAddr)
	}
}

func RunUDP() {

	addr, err := net.ResolveUDPAddr("udp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	// defer conn.Close()

	go func() {

		// server
		go func() {
			go handleUDPConnection(conn)
		}()

		// client
		go func() {
			addr, err := net.ResolveUDPAddr("udp", ":8081")
			if err != nil {
				log.Fatal(err)
			}

			conn, err := net.DialUDP("udp", nil, addr)
			// conn, err := net.ListenUDP("udp", addr)
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()

			message := "Hello\nUDP!\n"
			_, err = conn.Write([]byte(message))
			// _, err = conn.WriteToUDP([]byte(message), addr)
			if err != nil {
				log.Fatal(err)
			}

			buffer := make([]byte, 1024) // 고정 크기 버퍼
			n, _, _ := conn.ReadFromUDP(buffer)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Client: Server response:", string(buffer[:n]))
		}()
	}()
}
