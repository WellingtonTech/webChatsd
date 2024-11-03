package main

import (
	"fmt"
	"log"
	"net/http"
)

// define upgrader
// Read and Write buffer size
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})

	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Printf("Server start at localhos:8080")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
