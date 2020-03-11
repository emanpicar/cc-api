# CC API

[![Golang](https://golang.org/lib/godoc/images/go-logo-blue.svg)](https://golang.org/)

CC API is a simple microservice capable of handling REST API and validate credit cards via Luhn Algorithm

### Tech

CC API uses a number of open source projects to work properly:

* [Golang](https://golang.org/) - Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
* [gorilla/mux](https://github.com/gorilla/mux) - Package mux implements a request router and dispatcher.
* [Docker](https://www.docker.com/) - Securely build, share and run modern applications anywhere
* [jwt-go](https://github.com/dgrijalva/jwt-go) - A go (or 'golang' for search engine friendliness) implementation of JSON Web Tokens

### Installation

CC API requires [Docker](https://www.docker.com/) and [docker-compose](https://docs.docker.com/compose/) to run.

Install Docker and docker-compose to start the server
 - [Docker Desktop on Windows](https://docs.docker.com/docker-for-windows/install/)
 - [Docker on Linux](https://docs.docker.com/install/linux/docker-ce/centos/)
 - [Docker Desktop on MacOS](https://docs.docker.com/docker-for-mac/install/)
 - [Install docker-compose](https://docs.docker.com/compose/install/)

```sh
$ cd cc-api
$ docker-compose up
```
OR
```sh
$ cd cc-api
$ go run main.go #devmode
```

### Usage
    - POST "https://{HOST}:9988/api/authenticate"
        {
            "username": myuser,
            "password": mypass
        }
    - POST "https://{HOST}:9988/api/validateCards"
        [
            "4111111111111111",
        	"4111111111111",
        	"4012888888881881",
        	"378282246310005",
        	"6011111111111117",
        	"5105105105105100",
        	"5105 1051 0510 5106",
        	"9111111111111111"
        ]

### Todos

 - Write MORE Tests
 - Validate credentials against DB
 - Complete my unfinished JWT implementation
