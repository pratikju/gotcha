package main

import(
  "golang.org/x/net/websocket"
  "html/template"
  "net/http"
  "fmt"
  "os"
  "encoding/json"
  "flag"
)

var(
  host_address = flag.String("host","localhost:8000","providing the context")
  pwd, _        = os.Getwd()
  home_template = template.Must(template.ParseFiles(pwd + "/home.html"))
  Message       = websocket.Message
  ActiveClients = make(map[Client]int)     // map containing clients
)
//TODO Name to be included
type Client struct {
  websocket *websocket.Conn
  clientIP string
}

func init(){
  http.HandleFunc("/",home_handler)
  http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("."))))
  http.Handle("/websocket", websocket.Handler(SocketServer))
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

func home_handler(w http.ResponseWriter, r *http.Request){
  name := r.URL.Path[1:]
  if name == "" {
    name = "random person"
  }
  parsedUrl := map[string]string{"context": *host_address, "name": name }
  json, _ := json.Marshal(parsedUrl)
  err := home_template.Execute(w, string(json))
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func main(){
  flag.Parse()
  err := http.ListenAndServe(*host_address,nil)
  if err != nil {
    panic(err)
  }
}
