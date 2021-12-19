# Builder
FROM golang:1.17.3-alpine3.13 AS builderMonolith
WORKDIR /cont
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN apk --update add git make
RUN go build -o monolith ./cmd/main.go

FROM alpine:latest
RUN mkdir -p /app
WORKDIR /app

COPY --from=builderMonolith ./cont/monolith /app

CMD ["/app/monolith"]

