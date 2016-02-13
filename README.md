# go-chat

**go-chat** is a chat-application developed with the help of websockets, server side is implemented in golang and client side implemetation uses jQuery, bootstrap, emojify and jQuery-fileupload

## Installation

Assuming you have installed a recent version of
[Go](https://golang.org/doc/install), you can simply run

```
go get -u github.com/pratikju/go-chat
```

This will download Servidor to `$GOPATH/src/github.com/pratikju/go-chat`. From
  this directory run `go build` to create the `go-chat` binary.

## Usage

Start the server by executing `go-chat` binary. By default, server will listen to http://0.0.0.0:8000 for incoming requests.

```
go-chat -h
Usage of go-chat:
  -b string
    	listen on HOST (default "0.0.0.0")
  -p int
    	use PORT for HTTP (default 8000)
```
## Getting started

Connect to the server and start chatting with your friends. You can also share images, videos and lots more.
