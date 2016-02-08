package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

var (
	hostname      = flag.String("b", "localhost", "listen on HOST")
	port          = flag.Int("p", 8000, "use PORT for HTTP")
	// Message is websocket message encoder
	Message       = websocket.Message
	// ActiveClients is a map of websocket clients
	ActiveClients = make(map[Client]int) // map containing clients
)

// Client is a websocket client
type Client struct {
	websocket *websocket.Conn
	clientIP  string
}

func init() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/authorize_github", githubAuthorizationHandler)
	http.HandleFunc("/git_home", gitHomeHandler)
	http.HandleFunc("/authorize_google", googleAuthorizationHandler)
	http.HandleFunc("/google_home", googleHomeHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/uploads/", uploadViewHandler)
	http.Handle("/assets/", http.FileServer(http.Dir(".")))
	http.Handle("/websocket", websocket.Handler(socketServer))
}

func socketServer(ws *websocket.Conn) {
	var clientMessage string

	// cleanup on server side
	defer func() {
		if err := ws.Close(); err != nil {
			fmt.Println("Websocket could not be closed", err.Error())
		}
	}()

	clientIP := ws.Request().RemoteAddr
	newClient := Client{ws, clientIP}
	ActiveClients[newClient] = 0
	fmt.Println("Number of clients connected ...", len(ActiveClients))

	for {
		if err := Message.Receive(ws, &clientMessage); err != nil {
			delete(ActiveClients, newClient)
			fmt.Println("Number of clients still connected ...", len(ActiveClients))
			return
		}
		broadcastMessage(clientMessage)
	}
}

func broadcastMessage(clientMessage string) {
	for client := range ActiveClients {
		if err := Message.Send(client.websocket, clientMessage); err != nil {
			fmt.Println("Could not send message to ", client.clientIP, err.Error())
		}
	}
}

func httpListener(hostname string, port int) {
	host := fmt.Sprintf("%s:%d", hostname, port)

	if err := http.ListenAndServe(host, nil); err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()
	httpListener(*hostname, *port)
}
