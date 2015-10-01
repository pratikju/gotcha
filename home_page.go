package main

import(
  "html/template"
  "net/http"
  "fmt"
  "encoding/json"
)

const home_page = `
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8"/>
  <title> Go-Chat </title>
  <link rel="stylesheet" href="http://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css">
  <link rel="stylesheet" type="text/css" href="/assets/stylesheets/chat.css">
  <script src="//ajax.googleapis.com/ajax/libs/jquery/2.1.4/jquery.min.js"></script>
  <script src="http://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
  <script src="/assets/javascripts/chat.js"></script>
</head>
<body style="background-color: #FAFAFF;">
    <h1> Just Go-Chat</h1>
    <div id="data" style="display: none;">{{.}}</div>
    <div class="container-fluid">
      <div class="panel panel-primary">
        <div class="panel-heading">
          <span class="glyphicon glyphicon-comment"></span> Chat Box
          <div class="btn-group pull-right">
            <button type="button" class="btn btn-default btn-xs dropdown-toggle" data-toggle="dropdown">Options
              <span class="glyphicon glyphicon-chevron-down"></span>
            </button>
            <ul class="dropdown-menu slidedown">
              <li id="clear_chat"><a href="#"><span class="glyphicon glyphicon glyphicon-unchecked">
              </span>Clear Chat</a></li>
              <li id="leave_chat"><a href="#"><span class="glyphicon glyphicon glyphicon-remove-sign">
              </span>Leave Chat</a></li>
              <li id="join_chat"><a href=""><span class="glyphicon glyphicon-ok-sign">
              </span>Join Chat</a></li>
            </ul>
          </div>
        </div>
        <div id ="chat_box" class="panel-body msg_container_base"></div>

        <div class="panel-footer">
          <div class="input-group">
            <input id="chat_prompt" type="text" class="form-control input-sm" placeholder="Type your message here..." />
            <span class="input-group-btn">
              <button class="btn btn-warning btn-sm" id="send">Send</button>
            </span>
          </div>
        </div>
      </div>
    </div>
    <audio id="notify" preload="auto">
      <source src="http://demos.9lessons.info/notify/notify.ogg" type="audio/ogg">
      <source src="http://demos.9lessons.info/notify/notify.mp3" type="audio/mpeg">
      <source src="http://demos.9lessons.info/notify/notify.wav" type="audio/wav">
    </audio>
</body>
</html>
`

func home_handler(w http.ResponseWriter, r *http.Request){
  name := r.URL.Path[1:]
  if name == "" {
    name = "random person"
  }
  parsedUrl := map[string]string{"context": *host_address, "name": name }
  json, _ := json.Marshal(parsedUrl)
  home_template, error := template.New("webpage").Parse(home_page)
  if error != nil {
    fmt.Println("Couldn't parse home page!")
  }
  err := home_template.Execute(w, string(json))
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
