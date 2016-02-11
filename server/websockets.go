package server

import (
	"log"

	"golang.org/x/net/websocket"
)

type Client struct {
	websocket *websocket.Conn
	clientIP  string
}

var (
	// Message is websocket message encoder
	Message = websocket.Message
	// ActiveClients is a map of websocket clients
	ActiveClients = make(map[Client]int) // map containing clients
)

func socketHandler(ws *websocket.Conn) {
	var clientMessage string

	// cleanup on server side
	defer func() {
		if err := ws.Close(); err != nil {
			log.Println("Websocket could not be closed", err.Error())
		}
	}()

	clientIP := ws.Request().RemoteAddr
	newClient := Client{ws, clientIP}
	ActiveClients[newClient] = 0
	log.Println("Number of clients connected ...", len(ActiveClients))

	for {
		if err := Message.Receive(ws, &clientMessage); err != nil {
			delete(ActiveClients, newClient)
			log.Println("Number of clients still connected ...", len(ActiveClients))
			return
		}
		broadcastMessage(clientMessage)
	}
}

func broadcastMessage(clientMessage string) {
	for client := range ActiveClients {
		if err := Message.Send(client.websocket, clientMessage); err != nil {
			log.Println("Could not send message to ", client.clientIP, err.Error())
		}
	}
}
