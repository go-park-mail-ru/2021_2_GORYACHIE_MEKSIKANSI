# Builder
FROM golang:1.17.3-alpine3.13 AS builder

WORKDIR /
COPY . .
RUN apk update && apk upgrade && \
    apk --update add git make
COPY ../../../../go.mod .
RUN CGO_ENABLED=0 go get -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv
RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o engine ./cmd/main.go

FROM alpine:latest
RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app
WORKDIR /app

COPY --from=builder ./go/bin/dlv /app/dlv
COPY --from=builder ./engine /app
RUN mkdir internals
RUN cd internals
RUN mkdir microservice
RUN cd microservice
RUN mkdir cart
RUN cd cart
RUN mkdir config
RUN cd ..
RUN cd ..
RUN cd ..


CMD ["./dlv", "--listen=:40002", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/app/engine"]

#CMD /app/engine
