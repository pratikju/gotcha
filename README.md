[![Build Status](https://travis-ci.org/pratikju/gotcha.svg?branch=master)](https://travis-ci.org/pratikju/gotcha)

# gotcha

**gotcha** is a chat-application developed with the help of websockets, server side is implemented in golang and client side implemetation uses jQuery, bootstrap, emojify and jQuery-fileupload

## Installation

Assuming you have installed a recent version of
[Go](https://golang.org/doc/install), you can simply run

```
go get -u github.com/pratikju/gotcha
```

This will download gotcha to `$GOPATH/src/github.com/pratikju/gotcha`. From
  this directory run `go build` to create the `gotcha` binary.

## Usage

Start the server by executing `gotcha` binary. By default, server will listen to http://0.0.0.0:8000 for incoming requests.

```
gotcha -h
Usage of gotcha:
  -b string
    	listen on HOST (default "0.0.0.0")
  -p int
    	use PORT for HTTP (default 8000)
```
## Getting started

Connect to the server and start chatting with your friends. You can also share images, videos and lots more.

## Authentication

Oauth2 authentication is added at login. Till now, only google and github provider are included.
To add more providers,
 - Add another package in oauth folder.
 - Register the application at provider's developer console.
 - Set Client ID, Client secret, redirect url, endpoint, scope and profilesURL in config file in that package.
 - Then add the routes and corresponding handlers.

 ```
 	http.HandleFunc("/authorize_<provider>", <provider>AuthorizationHandler)
	http.HandleFunc("/<provider>_home", <provider>CallbackHandler)
 ```

 ```
 func <provider>AuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	url := <provider>.AuthConfig.AuthCodeURL("")
	http.Redirect(w, r, url, http.StatusFound)
}

func <provider>CallbackHandler(w http.ResponseWriter, r *http.Request) {
	config := <provider>.AuthConfig
	profilesURL := <provider>.ProfilesURL
	code := r.FormValue("code")
	handleCallback(w, r, config, profilesURL, code)
}
```

## License

MIT, see the [LICENSE](https://raw.githubusercontent.com/pratikju/gotcha/master/LICENSE.md) file.
