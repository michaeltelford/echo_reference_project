# bin

The bin directory contains binary executables used by Docker (in production) and docker-compose (in development). 

The binaries are all built for linux-amd64 e.g.

`GOOS=linux GOARCH=amd64 go build main.go`

Therefore the binaries aren't meant for local use (on a Mac for example). Instead, you should simply install the relavent libraries on your local development machine as needed. 
