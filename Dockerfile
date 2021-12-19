# Builder
FROM golang:1.17.3-alpine3.13 AS builderMonolith
WORKDIR /cont
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN apk --no-cache update && apk --no-cache upgrade && \
    apk --update --no-cache add git make
RUN go build -o monolith ./cmd/main.go

FROM alpine:latest
RUN apk --no-cache update && apk --no-cache upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app
WORKDIR /app

COPY --from=builderMonolith ./cont/monolith /app

CMD ["/app/monolith"]

