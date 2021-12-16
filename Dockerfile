# Builder
FROM golang:1.17.3-alpine3.13 AS builder

WORKDIR /
COPY . .
RUN apk update && apk upgrade && \
    apk --update add git make
RUN CGO_ENABLED=0 go get -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv
RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o engine ./cmd/main.go

FROM alpine:latest
RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app
WORKDIR /app

COPY --from=builder ./go/bin/dlv /app/dlv
COPY --from=builder ./engine /app
RUN mkdir config
RUN mkdir build
RUN cd build
RUN mkdir postgresql
RUN cd ..
RUN cd ..

CMD ["/app/engine"]

