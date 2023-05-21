package usecase

import (
	"HelloWorld/pkg/helloworld"
	"context"
)

type IHelloUsecase interface {
	SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloResponse, error)
}

type helloUsecase struct {
}

func NewHelloUsecase() IHelloUsecase {
	return &helloUsecase{}
}

func (h *helloUsecase) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {
	return &helloworld.HelloResponse{
		Message: "Hello " + in.Name,
	}, nil
}
