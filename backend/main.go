package main

import (
	"fmt"
	"net/http"

	"backend/pkg/websocket"
)

// define upgrader
// Read and Write buffer size
func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Printf("Server start at localhost:8080")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
