package server

import (
	"fmt"
	"log"
	"net/http"
)

//HTTPListener is a server
func HTTPListener(hostname string, port int) {
	host := fmt.Sprintf("%s:%d", hostname, port)
	log.Println("starting http server at", host)
	if err := http.ListenAndServe(host, nil); err != nil {
		log.Fatal(err)
	}
}
