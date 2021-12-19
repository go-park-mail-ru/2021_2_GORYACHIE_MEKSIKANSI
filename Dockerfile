# Builder
FROM golang:1.17.3-alpine3.13 AS builderMonolith
WORKDIR /cont
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN apk add --update-cache \
    apk update && apk upgrade && \
    apk --update add git make
RUN go build -o monolith ./cmd/main.go

FROM alpine:latest
RUN apk add --update-cache \
    apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app
WORKDIR /app

COPY --from=builderMonolith ./cont/monolith /app

CMD ["/app/monolith"]

