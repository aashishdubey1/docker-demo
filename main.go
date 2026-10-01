package main

import (
	"fmt"
	"net/http"
)

func handlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from the server")
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/ping", handlePing)
	mux.HandleFunc("/hello", handleHello)

	fmt.Println("server is running on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("error starting the server", err)
		return
	}
}
