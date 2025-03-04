package methods

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func RunJSONRPC() {

	arith := new(Arith)
	rpc.Register(arith)

	listener, err := net.Listen("tcp", ":8084")
	if err != nil {
		log.Fatal(err)
	}

	go func() {

		// server
		go func() {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal(err)
			}

			go rpc.ServeConn(conn)
		}()

		// client
		go func() {
			conn, err := rpc.Dial("tcp", ":8084")
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()

			args := Args{A: 7, B: 8}
			var reply int
			conn.Call("Arith.Multiply", &args, &reply)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println("Result:", reply)
		}()
	}()
}
