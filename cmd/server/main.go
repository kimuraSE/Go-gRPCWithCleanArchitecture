package main

import (
	"HelloWorld/internal/app/handler"
	"HelloWorld/internal/app/usecase"
	"HelloWorld/pkg/helloworld"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	helloworld.UnimplementedHelloServiceServer
}

func main() {

	helloUsecase := usecase.NewHelloUsecase()
	helloHandler := handler.NewHelloHandler(helloUsecase)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	helloworld.RegisterHelloServiceServer(s, helloHandler)
	log.Println("Server is running on port 50051")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
