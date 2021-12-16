# Builder
FROM golang:1.17.3-alpine3.13 AS builderRes
WORKDIR /cont

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN apk update && apk upgrade && \
    apk --update add git make
RUN go build -o engine ./internals/microservice/restaurant/cmd/main.go

FROM alpine:latest
RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app
WORKDIR /app

COPY --from=builderRes ./cont/engine /app
RUN mkdir internals
RUN cd internals
RUN mkdir microservice
RUN cd microservice
RUN mkdir restaurant
RUN cd restaurant
RUN mkdir config
RUN cd ..
RUN cd ..
RUN cd ..

CMD ["/app/engine"]

#CMD /app/engine
