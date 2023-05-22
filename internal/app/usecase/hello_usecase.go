package usecase

import (
	"HelloWorld/pkg/helloworld"
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

type IHelloUsecase interface {
	SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloResponse, error)
	Download(req *helloworld.DownloadRequest, stream helloworld.HelloService_DownloadServer) error
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

func (h *helloUsecase) Download(req *helloworld.DownloadRequest, stream helloworld.HelloService_DownloadServer) error {
	fmt.Println("Download")

	filename := req.GetFilename()
	path := "/home/kimuraryota/Document/dev/HelloWorld/pkg/storage/" + filename
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	buf := make([]byte, 5)
	for {
		n, err := file.Read(buf)
		if err != nil {
			return err
		}
		if n == 0 || err == io.EOF {
			break
		}
		res := &helloworld.DownloadResponse{
			Data: buf[:n],
		}
		sendErr := stream.Send(res)
		if sendErr != nil {
			return sendErr
		}
		time.Sleep(1 * time.Second)
	}

	return nil

}
