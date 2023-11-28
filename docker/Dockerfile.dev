ARG GO_VERSION=1.21
ARG ALPINE_VERSION=3.18

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION}

WORKDIR /go/src/plesio-server

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN apk update && apk add git
RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]