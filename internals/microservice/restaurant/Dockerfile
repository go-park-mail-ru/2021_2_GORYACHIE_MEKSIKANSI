# Builder
FROM golang:1.17.3-alpine3.13 AS builderRestaurant
WORKDIR /cont

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN apk update && apk upgrade && \
    apk --update add git make
RUN go build -o resService ./internals/microservice/restaurant/cmd/main.go

FROM alpine:latest
RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app
WORKDIR /app

COPY --from=builderRestaurant ./cont/resService /app

CMD ["/app/resService"]
