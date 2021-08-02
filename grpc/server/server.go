package server

import (
	"context"
	proto "github.com/how-to-go/grpc/generated/proto/v1"
)

type Server struct {
}

func (s *Server) Reverse(ctx context.Context, request *proto.ReverseRequest) (*proto.ReverseResponse, error) {
	runes := []rune(request.GetText())

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return &proto.ReverseResponse{
		ReversedText: string(runes),
	}, nil
}

func (s *Server) GetBar(ctx context.Context, request *proto.GetBarRequest) (*proto.GetBarResponse, error) {
	return &proto.GetBarResponse{Bar: &proto.Bar{
		Id:   42,
		Name: "Ultimate Question of Life, the Universe, and Everything",
	}},nil
}

// NewServer constructor
func NewServer() *Server {
	return &Server{}
}
