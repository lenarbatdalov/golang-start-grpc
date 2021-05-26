package main

import (
	"log"
	"net"

	"github.com/lenarbatdalov/grpcadder/pkg/adder"
	"github.com/lenarbatdalov/grpcadder/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	// grpc сервер
	s := grpc.NewServer()

	// переменная, которая реализует интерфейс сервера
	srv := &adder.GRPCServer{}

	// зарегистрирую adder сервер в качестве сервиса для grpc сервера
	api.RegisterAdderServer(s, srv)

	// создать слушателя, ктрый будет принимать запросы
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	// вызову у grpc сервера метод Server, ктрый принимает слушателя
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
