package methods

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, HTTP!")
}

func RunHTTP() {
	http.HandleFunc("/", handler)

	go func() {
		fmt.Println("Starting HTTP server on :8082")
		err := http.ListenAndServe(":8082", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
}
