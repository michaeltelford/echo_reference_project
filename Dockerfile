FROM alpine:3.8

WORKDIR /go/src/github.com/michaeltelford/echo_reference_project
COPY bin/api bin/api

CMD ["bin/api"]
