FROM golang:1.10

WORKDIR /go/src/github.com/michaeltelford/echo_reference_project
COPY . .

RUN go build main.go

CMD ["./main"]
