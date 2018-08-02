FROM golang:1.10

WORKDIR /go/src/github.com/michaeltelford/echo_reference_project
COPY bin/api bin/api

CMD ["bin/api"]
