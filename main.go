package main

import (
	"flag"

	"github.com/pratikju/go-chat/server"
	"github.com/pratikju/go-chat/session"
)

var (
	hostname = flag.String("b", "0.0.0.0", "hostname to be used")
	port     = flag.Int("p", 8000, "port on which server will listen")
)

func main() {
	flag.Parse()
	session.Init()
	server.AttachHandlers()
	server.ListenHTTP(*hostname, *port, nil)
}
