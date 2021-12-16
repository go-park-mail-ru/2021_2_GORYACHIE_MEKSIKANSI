# Builder
FROM golang:1.17.3-alpine3.13 AS builder

WORKDIR /cont
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN apk update && apk upgrade && \
    apk --update add git make
RUN go build -o engine ./internals/microservice/promocode/cmd/main.go

FROM alpine:latest
RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app
WORKDIR /app

COPY --from=builder ./cont/engine /app
RUN mkdir internals
RUN cd internals
RUN mkdir microservice
RUN cd microservice
RUN mkdir promocode
RUN cd promocode
RUN mkdir config
RUN cd ..
RUN cd ..
RUN cd ..


CMD ["/app/engine"]

#CMD /app/engine
