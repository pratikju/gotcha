# go-chat

**go-chat** is a chat-application developed with the help of websockets, server side is implemented in golang and client side implemetation uses jQuery, bootstrap, emojify and jQuery-fileupload

## Dependencies

Only websocket package needs to be imported explicitly.

```
go get "golang.org/x/net/websocket"

```
## Usage

Build the go project and execute. By default, hostname: **localhost** and port: **8000** . Same can be changed by providing hostname and port at run-time.

```
  go build  &&
  ./go-chat -b <hostname> -p <port>
```
## Getting started

Connect to the server and start chatting with your friends. You can also share images, videos and lots more.
