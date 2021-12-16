# Builder
FROM golang:1.17.3-alpine3.13 AS builderAuth
WORKDIR /cont

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN apk update && apk upgrade && \
    apk --update add git make
RUN go build -o engine ./internals/microservice/authorization/cmd/main.go

FROM alpine:latest
RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app
WORKDIR /app

COPY --from=builderAuth ./cont/engine /app
RUN mkdir internals
RUN cd internals
RUN mkdir microservice
RUN cd microservice
RUN mkdir authorization
RUN cd authorization
RUN mkdir config
RUN cd ..
RUN cd ..
RUN cd ..

CMD ["/app/engine"]

#CMD /app/engine
