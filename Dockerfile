FROM golang:alpine as base

WORKDIR /usr/src/users-api

ENV GO111MODULE=on
ENV GOOS="linux"
ENV CGO_ENABLED=0

# System dependencies

RUN apk --no-cache add build-base ca-certificates

## Dev (Hot reload and debugger)

FROM base as dev

WORKDIR /usr/src/users-api

RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go get -u github.com/cosmtrek/air

EXPOSE 8080
EXPOSE 2345

ENTRYPOINT [ "air" ]

# Builder

FROM base as builder

WORKDIR /usr/src/users-api

COPY . .

RUN go mod download && go mod verify

RUN go build -o users-api -a .

# Production

FROM alpine:latest

#Copy executable from builder
COPY --from=builder /usr/src/users-api/users-api /usr/local/bin/users-api

EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/users-api"]
