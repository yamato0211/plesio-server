ARG GO_VERSION=1.21
ARG ALPINE_VERSION=3.18

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as go-builder

WORKDIR /go/src/plesio-server

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o api cmd/main.go


FROM alpine:${ALPINE_VERSION}

WORKDIR /usr/src/plesio-server

COPY --from=go-builder /go/src/plesio-server/api api
COPY --from=go-builder /go/src/plesio-server/api .env
RUN chmod +x "/usr/src/plesio-server/api"

ENTRYPOINT ["/usr/src/plesio-server/api"]