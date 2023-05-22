package handler

import (
	"HelloWorld/internal/app/usecase"
	"HelloWorld/pkg/helloworld"
	"context"
)

type IHelloHandler interface {
	SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloResponse, error)
	Download(req *helloworld.DownloadRequest, stream helloworld.HelloService_DownloadServer) error
}

type helloHandler struct {
	helloworld.UnimplementedHelloServiceServer
	hu usecase.IHelloUsecase
}

func NewHelloHandler(hu usecase.IHelloUsecase) helloworld.HelloServiceServer {
	return &helloHandler{
		hu: hu,
	}
}

func (h *helloHandler) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {
	res, err := h.hu.SayHello(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (h *helloHandler) Download(req *helloworld.DownloadRequest, stream helloworld.HelloService_DownloadServer) error {
	err := h.hu.Download(req, stream)
	if err != nil {
		return err
	}
	return nil
}
