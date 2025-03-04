package methods

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, HTTP2!")
}

func RunHTTP2() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler2)

	server := &http.Server{
		Addr:    ":8083",
		Handler: mux,
	}
	err := http2.ConfigureServer(server, &http2.Server{})
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		fmt.Println("Starting HTTP/2 server on :8083")
		err = server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
}
