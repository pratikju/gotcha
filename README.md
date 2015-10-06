# go-chat

**go-chat** is a chat-application in go with usage of websockets.

## Dependencies

Only websocket package needs to be imported explicitly.

```
go get "golang.org/x/net/websocket"

```
## Usage

Build the project and execute. By default, host is **localhost:8000** . Same can be changed by providing host at run-time.

```
  go build  &&
  ./go-chat -host <hostname:port>
```
## Getting started

Connect to the server by providing your name after the host and enjoy chatting with your friends..

```
<hostname:port>/<your-name>
```
