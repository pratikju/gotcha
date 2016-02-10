package server

import (
	"fmt"
	"log"
	"net/http"
)

// ListenHTTP starts http server at given hostname and port
func ListenHTTP(hostname string, port int) {
	host := fmt.Sprintf("%s:%d", hostname, port)
	log.Println("starting http server at", host)
	if err := http.ListenAndServe(host, nil); err != nil {
		log.Fatal(err)
	}
}
