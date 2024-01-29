package main

import (
	"net"
	"net/http"
)

func main() {
	// On écoute sur le port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	// On démarre le serveur HTTP
	http.Serve(listener, nil)

}