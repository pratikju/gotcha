package main

import(
  "golang.org/x/net/websocket"
  "net/http"
  "fmt"
  "flag"
)

var(
  hostname = flag.String("b", "localhost", "listen on HOST")
	port = flag.Int("p", 8000, "use PORT for HTTP")
  Message       = websocket.Message
  ActiveClients = make(map[Client]int)  // map containing clients
)
//TODO Name to be included
type Client struct {
  websocket *websocket.Conn
  clientIP string
}

func init(){
  http.HandleFunc("/home",home_handler)
  http.HandleFunc("/login",login_handler)
  http.HandleFunc("/",redirect_handler)
  http.HandleFunc("/upload", upload_handler)
  http.Handle("/assets/", http.FileServer(http.Dir(".")))
  http.Handle("/websocket", websocket.Handler(SocketServer))
  http.HandleFunc("/uploads/", upload_view_Handler)
}

func SocketServer(ws *websocket.Conn) {
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

func broadcastMessage(clientMessage string){
  for client, _ := range ActiveClients {
    if err := Message.Send(client.websocket, clientMessage); err != nil {
      fmt.Println("Could not send message to ", client.clientIP, err.Error())
    }
  }
}

func HTTPListener(hostname string, port int) {
	host := fmt.Sprintf("%s:%d", hostname, port)

	if err := http.ListenAndServe(host, nil); err != nil {
    panic(err)
  }
}

func main(){
  flag.Parse()
  HTTPListener(*hostname, *port)
}
