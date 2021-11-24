package main

import (
	"google.golang.org/grpc"
	"net"
)

func main() {
	listen, errListen := net.Listen("tcp", ":8081")
	if errListen != nil {
		println("GG")
	}
	server := grpc.NewServer()


	println("Listen in 127.0.0.1:8081")
	errServ := server.Serve(listen)
	if errServ != nil {
		println("GG Serve")
		return
	}

}
